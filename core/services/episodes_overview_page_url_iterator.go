package services

import (
	"github.com/sarulabs/di"
	"rkl.io/kika-downloader/core/crawler"
	"rkl.io/kika-downloader/core/http"
)

// AssignEpisodesOverviewUrlIterator assign http client service
func AssignEpisodesOverviewUrlIterator(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "episodes_overview_url_iterator",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			service, err := ctx.SafeGet("http_client")
			if err != nil {
				return nil, err
			}
			httpClient := service.(http.ClientInterface)

			episodesOverviewCrawler := crawler.NewEpisodesOverviewUrlIterator(
				httpClient,
				ctx.Get("xpath_episodes_overview_page_page_items").(string),
			)

			return episodesOverviewCrawler.(crawler.IteratorInterface), nil
		},
	})

	return nil
}
