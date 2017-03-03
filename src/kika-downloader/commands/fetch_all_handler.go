package commands

import (
	"fmt"
	"kika-downloader/contract"
	"kika-downloader/crawler"
	"log"
	"net/url"
)

// ProgressDTO DTO to export progress
type ProgressDTO struct {
	Percentage   string
	SeriesTitle  string
	EpisodeTitle string
}

type fetchAllHandler struct {
	command *FetchAll

	episodesPageIterator crawler.IteratorInterface
	pageItemsIterator    crawler.IteratorInterface
	videoExtractor       contract.VideoExtractorInterface
	videoDownloader      contract.VideoDownloaderInterface
}

// NewFetchAllHandler return new fetch all command handler
func NewFetchAllHandler(
	episodesPageIterator,
	pageItemsIterator crawler.IteratorInterface,
	videoExtractor contract.VideoExtractorInterface,
	videoDownloader contract.VideoDownloaderInterface,

) contract.CommandHandlerInterface {
	return &fetchAllHandler{
		episodesPageIterator: episodesPageIterator,
		pageItemsIterator:    pageItemsIterator,
		videoExtractor:       videoExtractor,
		videoDownloader:      videoDownloader,
	}
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

func (h *fetchAllHandler) handle(command *FetchAll) (interface{}, error) {
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

			// download video
			if err := h.downloadVideo(video, command.GetOutputDir()); err != nil {
				return nil, err
			}
		}
	}

	return nil, nil
}

func (h *fetchAllHandler) downloadVideo(video contract.VideoInterface, outputDir string) error {
	progressChannel, err := h.videoDownloader.Download(video, outputDir)
	if err != nil {
		return err
	}

	// progress handling
	for p := range progressChannel {
		if err := h.videoDownloader.GetLastError(); err != nil {
			return err
		}

		//h.progress <- ProgressDTO{
		//	p.GetPercentage(),
		//	video.GetSeriesTitle(),
		//	video.GetEpisodeTitle(),
		//}

		fmt.Printf("\r[%s %%] %s - %s\n", p.GetPercentage(), video.GetSeriesTitle(), video.GetEpisodeTitle())
	}

	return nil
}
