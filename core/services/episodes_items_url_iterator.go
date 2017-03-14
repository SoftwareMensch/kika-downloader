package services

import (
	"github.com/sarulabs/di"
	"rkl.io/kika-downloader/core/crawler"
	"rkl.io/kika-downloader/core/http"
)

// AssignEpisodesOverviewItemsUrlIterator assign http client service
func AssignEpisodesOverviewItemsUrlIterator(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "episodes_items_url_iterator",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			service, err := ctx.SafeGet("http_client")
			if err != nil {
				return nil, err
			}
			httpClient := service.(http.ClientInterface)

			itemsCrawler := crawler.NewEpisodesItemsUrlIterator(
				httpClient,
				ctx.Get("xpath_episodes_items").(string),
			)

			return itemsCrawler.(crawler.IteratorInterface), nil
		},
	})

	return nil
}
