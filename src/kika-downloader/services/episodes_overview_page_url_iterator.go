package services

import (
	"github.com/sarulabs/di"
	"kika-downloader/crawler"
	"kika-downloader/http"
	"net/url"
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

			episodesOverviewCrawler.SetCrawlingURL(ctx.Get("episodes_overview_url").(*url.URL))

			return episodesOverviewCrawler.(crawler.IteratorInterface), nil
		},
	})

	return nil
}
