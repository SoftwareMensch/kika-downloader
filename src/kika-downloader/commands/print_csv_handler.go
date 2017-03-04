package commands

import (
	"fmt"
	"kika-downloader/contract"
	"kika-downloader/crawler"
	"kika-downloader/dto"
	"net/url"
	"strings"
)

type printCsvHandler struct {
	command *PrintCsv

	episodesPageIterator crawler.IteratorInterface
	pageItemsIterator    crawler.IteratorInterface
	videoExtractor       contract.VideoExtractorInterface
}

// NewPrintCsvHandler return new fetch all command handler
func NewPrintCsvHandler(
	episodesPageIterator,
	pageItemsIterator crawler.IteratorInterface,
	videoExtractor contract.VideoExtractorInterface,

) contract.CommandHandlerInterface {
	return &printCsvHandler{
		episodesPageIterator: episodesPageIterator,
		pageItemsIterator:    pageItemsIterator,
		videoExtractor:       videoExtractor,
	}
}

// Handle handle command
func (h *printCsvHandler) Handle(command interface{}) (interface{}, error) {
	switch t := command.(type) {
	case *PrintCsv:
		return h.handle(command.(*PrintCsv))
	default:
		return nil, fmt.Errorf("cannot handle command of type \"%s\"", t)
	}
}

func (h *printCsvHandler) handle(command *PrintCsv) (interface{}, error) {
	overviewURL := command.GetOverviewUrl()
	h.episodesPageIterator.SetCrawlingURL(overviewURL)

	fmt.Printf("No.;Series Title;Episode Title;Episode Description;Video URL\n")

	for rawPageURL := range h.episodesPageIterator.Run() {
		pageURL, err := url.Parse(rawPageURL)
		if err != nil {
			return nil, err
		}

		// iterate urls for the videos
		h.pageItemsIterator.SetCrawlingURL(pageURL)

		for rawItemURL := range h.pageItemsIterator.Run() {
			itemURL, err := url.Parse(rawItemURL)
			if err != nil {
				return nil, err
			}

			// extract video from item url
			video, err := h.videoExtractor.ExtractVideoFromURL(itemURL.String())
			if err != nil {
				return nil, err
			}

			videoDTO := dto.CsvLineFromVideo(video)

			fmt.Printf(
				"%02d;%s;%s;%s;%s\n",
				videoDTO.GetEpisodeNr(),
				strings.Replace(videoDTO.GetSeriesTitle(), ";", "", -1),
				strings.Replace(videoDTO.GetEpisodeTitle(), ";", "", -1),
				strings.Replace(videoDTO.GetEpisodeDescription(), ";", "", -1),
				videoDTO.GetDownloadUrl(),
			)
		}
	}

	return nil, nil
}
