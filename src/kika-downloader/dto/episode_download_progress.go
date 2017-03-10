package dto

// EpisodeDownloadProgress DTO to export progress
type EpisodeDownloadProgress struct {
	percentage   string
	seriesTitle  string
	episodeTitle string
}

// NewEpisodeDownloadProgress
func NewEpisodeDownloadProgress(percentage, seriesTitle, episodeTitle string) EpisodeDownloadProgress {
	return EpisodeDownloadProgress{percentage, seriesTitle, episodeTitle}
}

// GetPercentage
func (p EpisodeDownloadProgress) GetPercentage() string {
	return p.percentage
}

// GetSeriesTitle
func (p EpisodeDownloadProgress) GetSeriesTitle() string {
	return p.seriesTitle
}

// GetEpisodeTitle
func (p EpisodeDownloadProgress) GetEpisodeTitle() string {
	return p.episodeTitle
}
