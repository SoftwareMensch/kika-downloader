package contract

import (
	"net/url"
	"rkl.io/kika-downloader/core/vo"
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
