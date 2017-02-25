package model

import (
	"kika-downloader/contract"
	"kika-downloader/vo"
	"net/url"
)

type video struct {
	abstractEntity

	resolution         vo.Resolution
	seriesTitle        string
	episodeTitle       string
	episodeDescription string
	languageCode       string
	originURL          *url.URL
	episodeNumber      int
	fileSize           int
}

// NewVideo instantiate new video
func NewVideo(
	id string,
	seriesTitle string,
	episodeTitle string,
	episodeDescription string,
	languageCode string,
	episodeNumber int,
	resolution vo.Resolution,
	originURL *url.URL,
	fileSize int,

) contract.VideoInterface {
	v := &video{
		seriesTitle:        seriesTitle,
		episodeTitle:       episodeTitle,
		episodeDescription: episodeDescription,
		languageCode:       languageCode,
		episodeNumber:      episodeNumber,
		resolution:         resolution,
		originURL:          originURL,
		fileSize:           fileSize,
	}

	v.abstractEntity.id = id

	return v
}

// GetSeriesTitle return series title
func (v *video) GetSeriesTitle() string {
	return v.seriesTitle
}

// GetEpisodeDescription return series title
func (v *video) GetEpisodeDescription() string {
	return v.episodeDescription
}

// GetEpisodeTitle return episodes title
func (v *video) GetEpisodeTitle() string {
	return v.episodeTitle
}

// GetLanguageCode return language code
func (v *video) GetLanguageCode() string {
	return v.languageCode
}

// GetVideoOriginURL return origin video url
func (v *video) GetVideoOriginURL() *url.URL {
	return v.originURL
}

// GetVideoResolution get video resolution
func (v *video) GetVideoResolution() vo.Resolution {
	return v.resolution
}

// GetEpisodeNumber get no. of episode
func (v *video) GetEpisodeNumber() int {
	return v.episodeNumber
}

// GetFileSize get no. of episode
func (v *video) GetFileSize() int {
	return v.fileSize
}
