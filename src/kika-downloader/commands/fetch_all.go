package commands

import (
	"net/url"
)

// FetchAll fetch all command
type FetchAll struct {
	overviewURL *url.URL
	outputDir   string
}

// NewFetchAll new fetch all command
func NewFetchAll(entryURL *url.URL, outputDir string) *FetchAll {
	return &FetchAll{entryURL, outputDir}
}

// GetOverviewURL return url of entry point
func (c *FetchAll) GetOverviewURL() *url.URL {
	return c.overviewURL
}

// GetOutputDir return path of output dir
func (c *FetchAll) GetOutputDir() string {
	return c.outputDir
}
