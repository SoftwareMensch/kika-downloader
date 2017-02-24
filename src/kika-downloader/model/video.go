package model

import "kika-downloader/vo"

type video struct {
	abstractEntity

	resolution vo.Resolution
	seriesTitle string
	episodeTitle string
	languageCode string
	originURL string
	episodeNumber int
}

func (v *video) GetSeriesTitle() string {
	return v.seriesTitle
}

func (v *video) GetEpisodeTitle() string {
	return v.episodeTitle
}

func (v *video) GetLanguageCode() string {
	return v.languageCode
}

func (v *video) GetVideoOriginURL() string {
	return v.originURL
}

func (v *video) GetVideoResolution() vo.Resolution {
	return v.resolution
}

func (v *video) GetEpisodeNumber() int {
	return v.episodeNumber
}
