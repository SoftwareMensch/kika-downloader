package commands

import (
	"kika-downloader/crawler"
	"fmt"
	"kika-downloader/contract"
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
	return nil, nil
}
