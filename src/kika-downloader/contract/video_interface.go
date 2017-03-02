package contract

import (
	"kika-downloader/vo"
	"net/url"
)

type VideoInterface interface {
	EntityInterface

	GetSeriesTitle() string
	GetEpisodeTitle() string
	GetEpisodeDescription() string
	GetLanguageCode() string
	GetVideoOriginURL() *url.URL
	GetVideoResolution() vo.Resolution
	GetEpisodeNumber() int
	GetFileSize() int64
}
