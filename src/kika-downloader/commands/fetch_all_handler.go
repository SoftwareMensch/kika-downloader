package commands

import (
	"fmt"
	"kika-downloader/contract"
)

type fetchAllHandler struct {
	command *FetchAll
}

// NewFetchAllHandler return new fetch all command handler
func NewFetchAllHandler() contract.CommandHandlerInterface {
	return &fetchAllHandler{}
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
	return nil, nil
}
