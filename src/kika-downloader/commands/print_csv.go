package commands

import "net/url"

// PrintCsv print csv command
type PrintCsv struct {
	overviewURL *url.URL
}

// NewPrintCsvCommand return print csv command
func NewPrintCsvCommand(overviewURL *url.URL) *PrintCsv {
	return &PrintCsv{overviewURL}
}

// GetOverviewUrl get overview url
func (c *PrintCsv) GetOverviewUrl() *url.URL {
	return c.overviewURL
}
