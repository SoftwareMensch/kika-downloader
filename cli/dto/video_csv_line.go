package dto

import "rkl.io/kika-downloader/core/contract"

type videoCsvLine struct {
	seriesTitle        string
	episodeTitle       string
	episodeDescription string
	downloadURL        string
	episodeNo          int
}

// CsvLineFromVideo create a csv line from a video
func CsvLineFromVideo(video contract.VideoInterface) videoCsvLine {
	return videoCsvLine{
		seriesTitle:        video.GetSeriesTitle(),
		episodeTitle:       video.GetEpisodeTitle(),
		episodeDescription: video.GetEpisodeDescription(),
		downloadURL:        video.GetVideoOriginURL().String(),
		episodeNo:          video.GetEpisodeNumber(),
	}
}

// GetSeriesTitle
func (l videoCsvLine) GetSeriesTitle() string {
	return l.seriesTitle
}

// GetEpisodeTitle
func (l videoCsvLine) GetEpisodeTitle() string {
	return l.episodeTitle
}

// GetEpisodeDescription
func (l videoCsvLine) GetEpisodeDescription() string {
	return l.episodeDescription
}

// GetDownloadUrl
func (l videoCsvLine) GetDownloadUrl() string {
	return l.downloadURL
}

// GetEpisodeNr
func (l videoCsvLine) GetEpisodeNr() int {
	return l.episodeNo
}
