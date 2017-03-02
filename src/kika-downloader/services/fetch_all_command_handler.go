package services

import (
	"github.com/sarulabs/di"
	"kika-downloader/commands"
	"kika-downloader/contract"
	"kika-downloader/http"
)

// AssignEpisodesOverviewItemsUrlIterator assign http client service
func AssignFetchAllCommandHandler(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "command_handler.fetch_all",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			service, err := ctx.SafeGet("http_client")
			if err != nil {
				return nil, err
			}
			httpClient := service.(http.ClientInterface)

			return commands.NewFetchAllHandler(httpClient).(contract.CommandHandlerInterface), nil
		},
	})

	return nil
}
