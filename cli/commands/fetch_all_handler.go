package commands

import (
	"fmt"
	"log"
	"net/url"
	"os"
	cliContract "rkl.io/kika-downloader/cli/contract"
	"rkl.io/kika-downloader/cli/dto"
	coreContract "rkl.io/kika-downloader/core/contract"
	"rkl.io/kika-downloader/core/crawler"
	"runtime"
	"sync"
)

type fetchAllHandler struct {
	abstractHandler

	command *FetchAll

	wg sync.WaitGroup

	maxSimultaneousDownloads     int
	currentSimultaneousDownloads int

	episodesPageIterator crawler.IteratorInterface
	pageItemsIterator    crawler.IteratorInterface
	videoExtractor       coreContract.VideoExtractorInterface
	videoDownloader      coreContract.VideoDownloaderInterface
}

// NewFetchAllHandler return new fetch all command handler
func NewFetchAllHandler(
	episodesPageIterator,
	pageItemsIterator crawler.IteratorInterface,
	videoExtractor coreContract.VideoExtractorInterface,
	videoDownloader coreContract.VideoDownloaderInterface,

) cliContract.CommandHandlerInterface {
	handler := &fetchAllHandler{
		episodesPageIterator: episodesPageIterator,
		pageItemsIterator:    pageItemsIterator,
		videoExtractor:       videoExtractor,
		videoDownloader:      videoDownloader,
	}

	handler.currentSimultaneousDownloads = 0
	handler.dtoOutputChannel = make(chan interface{})

	runtime.SetFinalizer(handler, func(h *fetchAllHandler) {
		if h.dtoOutputChannel != nil {
			close(h.dtoOutputChannel)
		}
	})

	return handler
}

// Handle handle command
func (h *fetchAllHandler) Handle(command interface{}) (interface{}, error) {
	switch t := command.(type) {
	case *FetchAll:
		return h.handle(command.(*FetchAll))
	default:
		return nil, fmt.Errorf("cannot handle command of type \"%s\"", t)
	}
}

// GetDtoOutputChannel get output channel
func (h *fetchAllHandler) GetDtoOutputChannel() chan interface{} {
	return h.dtoOutputChannel
}

func (h *fetchAllHandler) handle(command *FetchAll) (interface{}, error) {
	h.maxSimultaneousDownloads = command.GetMaxSimultaneousDownloads()

	overviewURL := command.GetOverviewURL()
	h.episodesPageIterator.SetCrawlingURL(overviewURL)

	log.Printf("fetch-all command started for => %s", overviewURL.String())

	for rawPageURL := range h.episodesPageIterator.Run() {
		pageURL, err := url.Parse(rawPageURL)
		if err != nil {
			return nil, err
		}

		// iterate urls for the videos
		h.pageItemsIterator.SetCrawlingURL(pageURL)

		for rawItemURL := range h.pageItemsIterator.Run() {
			// notify about download errors and go on
			if lastDownloaderError := h.videoDownloader.GetLastError(); lastDownloaderError != nil {
				log.Printf("[I] %s", lastDownloaderError.Error())
				h.videoDownloader.ResetLastError()
			}

			itemURL, err := url.Parse(rawItemURL)
			if err != nil {
				return nil, err
			}

			// extract video from item url
			video, err := h.videoExtractor.ExtractVideoFromURL(itemURL.String())
			if err != nil {
				return nil, err
			}

			if h.currentSimultaneousDownloads > h.maxSimultaneousDownloads-1 {
				h.wg.Wait()
				h.currentSimultaneousDownloads = 0
			}

			// download video
			go func() {
				h.currentSimultaneousDownloads++
				h.wg.Add(1)

				defer h.wg.Done()

				if err := h.downloadVideo(video, command.GetOutputDir()); err != nil {
					fmt.Fprintf(os.Stderr, "[E] %s\n", err.Error())
				}
			}()
		}
	}

	h.wg.Wait()

	return nil, nil
}

func (h *fetchAllHandler) downloadVideo(video coreContract.VideoInterface, outputDir string) error {
	progressChannel, err := h.videoDownloader.Download(video, outputDir)
	if err != nil {
		return err
	}

	// progress handling
	for p := range progressChannel {
		if err := h.videoDownloader.GetLastError(); err != nil {
			return err
		}

		progressDto := dto.NewEpisodeDownloadProgress(
			p.GetPercentage(),
			video.GetSeriesTitle(),
			video.GetEpisodeTitle(),
		)

		h.dtoOutputChannel <- progressDto
	}

	return nil
}
