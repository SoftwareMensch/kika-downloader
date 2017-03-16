package commands

import (
	"net/url"
)

// FetchAll fetch all command
type FetchAll struct {
	overviewURL              *url.URL
	outputDir                string
	maxSimultaneousDownloads int
}

// NewFetchAll new fetch all command
func NewFetchAll(entryURL *url.URL, outputDir string, maxSimultaneousDownloads int) *FetchAll {
	return &FetchAll{entryURL, outputDir, maxSimultaneousDownloads}
}

// GetOverviewURL return url of entry point
func (c *FetchAll) GetOverviewURL() *url.URL {
	return c.overviewURL
}

// GetOutputDir return path of output dir
func (c *FetchAll) GetOutputDir() string {
	return c.outputDir
}

// GetMaxSimultaneousDownloads return path of output dir
func (c *FetchAll) GetMaxSimultaneousDownloads() int {
	return c.maxSimultaneousDownloads
}
