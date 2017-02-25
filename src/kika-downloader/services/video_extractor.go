package services

import (
	"github.com/sarulabs/di"
	"kika-downloader/http"

	"kika-downloader/contract"
	"kika-downloader/extractor"
)

// AssignExtractor service to extract video meta data
func AssignExtractor(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "video_extractor",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			service, err := ctx.SafeGet("http_client")
			if err != nil {
				return nil, err
			}
			httpClient := service.(http.ClientInterface)

			xPathVideoPageVideoTags, err := ctx.SafeGet("xpath_video_page_video_tags")
			if err != nil {
				return nil, err
			}

			xPathVideoPageXmlDataTag, err := ctx.SafeGet("xpath_video_page_xml_data_tag")
			if err != nil {
				return nil, err
			}

			xPathXmlSeriesTitle, err := ctx.SafeGet("xpath_xml_series_title")
			if err != nil {
				return nil, err
			}

			xPathXmlEpisodeTitle, err := ctx.SafeGet("xpath_xml_episode_title")
			if err != nil {
				return nil, err
			}

			xPathXmlEpisodeLanguageCode, err := ctx.SafeGet("xpath_xml_episode_language")
			if err != nil {
				return nil, err
			}

			xPathXmlEpisodeDescription, err := ctx.SafeGet("xpath_xml_episode_description")
			if err != nil {
				return nil, err
			}

			xPathXmlAssets, err := ctx.SafeGet("xpath_xml_assets")
			if err != nil {
				return nil, err
			}

			regExpVideoId, err := ctx.SafeGet("regexp_video_id")
			if err != nil {
				return nil, err
			}

			regExpXmlDataUrl, err := ctx.SafeGet("regexp_xml_data_url")
			if err != nil {
				return nil, err
			}

			videoExtractor := extractor.NewXmlVideoExtractor(
				httpClient,
				xPathVideoPageVideoTags.(string),
				xPathVideoPageXmlDataTag.(string),
				xPathXmlSeriesTitle.(string),
				xPathXmlEpisodeTitle.(string),
				xPathXmlEpisodeLanguageCode.(string),
				xPathXmlEpisodeDescription.(string),
				xPathXmlAssets.(string),
				regExpVideoId.(string),
				regExpXmlDataUrl.(string),
			)

			return videoExtractor.(contract.VideoExtractorInterface), nil
		},
	})

	return nil
}
