package contract

import "kika-downloader/vo"

type VideoInterface interface {
	GetSeriesTitle() string
	GetEpisodeTitle() string
	GetLanguageCode() string
	GetVideoOriginURL() string
	GetVideoResolution() vo.Resolution
	GetEpisodeNumber() int
}