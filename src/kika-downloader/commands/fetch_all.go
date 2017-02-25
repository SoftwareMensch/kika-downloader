package commands

import (
	"net/url"
)

// FetchAll fetch all command
type FetchAll struct {
	entryURL *url.URL
}

// NewFetchAll new fetch all command
func NewFetchAll(entryURL *url.URL) *FetchAll {
	return &FetchAll{}
}

// GetEntryURL return url of entry point
func (c *FetchAll) GetEntryURL() *url.URL {
	return c.entryURL
}
