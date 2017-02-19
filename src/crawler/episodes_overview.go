package crawler

import "net/url"

// EpisodesOverview crawler for the episodes overview pages
type EpisodesOverview struct {
	entryURL *url.URL
}

// NewEpisodesOverview return new *EpisodesOverview
func NewEpisodesOverview(rawEntryURL string) (*EpisodesOverview, error) {
	parsedURL, err := url.Parse(rawEntryURL)
	if err != nil {
		return nil, err
	}

	return &EpisodesOverview{
		entryURL: parsedURL,
	}, nil
}
