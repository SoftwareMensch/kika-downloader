package services

import (
	"github.com/sarulabs/di"
	"kika-downloader/http"

	"kika-downloader/extractor"
	"kika-downloader/contract"
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
				regExpVideoId.(string),
				regExpXmlDataUrl.(string),
			)

			return videoExtractor.(contract.VideoExtractorInterface), nil
		},
	})

	return nil
}
