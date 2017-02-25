package services

import (
	"github.com/sarulabs/di"
	"kika-downloader/contract"
	"kika-downloader/downloader"
	"kika-downloader/http"
)

// AssignVideoDownloader assign http client service
func AssignVideoDownloader(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "video_downloader",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			service, err := ctx.SafeGet("http_client")
			if err != nil {
				return nil, err
			}

			httpClient := service.(http.ClientInterface)

			return downloader.NewVideoDownloader(httpClient).(contract.VideoDownloaderInterface), nil
		},
	})

	return nil
}
