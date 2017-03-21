package dto

import "rkl.io/kika-downloader/core/contract"

type Video struct {
	SeriesTitle        string `json:"series_title"`
	EpisodeTitle       string `json:"episode_title"`
	EpisodeDescription string `json:"episode_description"`
	DownloadURL        string `json:"download_url"`
	EpisodeNo          int    `json:"episode_no"`
}

// NewVideoDtoFromModel new from model
func NewVideoDtoFromModel(model contract.VideoInterface) Video {
	return Video{
		SeriesTitle:        model.GetSeriesTitle(),
		EpisodeTitle:       model.GetEpisodeTitle(),
		EpisodeDescription: model.GetEpisodeDescription(),
		DownloadURL:        model.GetVideoOriginURL().String(),
		EpisodeNo:          model.GetEpisodeNumber(),
	}
}
