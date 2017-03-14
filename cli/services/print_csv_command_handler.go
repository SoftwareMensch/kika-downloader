package services

import (
	"github.com/sarulabs/di"
	"rkl.io/kika-downloader/cli/commands"
	cliContract "rkl.io/kika-downloader/cli/contract"
	coreContract "rkl.io/kika-downloader/core/contract"
	"rkl.io/kika-downloader/core/crawler"
)

// AssignEpisodesOverviewItemsUrlIterator
func AssignPrintCsvCommandHandler(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "command_handler.print_csv",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			service, err := ctx.SafeGet("episodes_overview_url_iterator")
			if err != nil {
				return nil, err
			}
			episodesPageIterator := service.(crawler.IteratorInterface)

			service, err = ctx.SafeGet("episodes_items_url_iterator")
			if err != nil {
				return nil, err
			}
			pageItemsIterator := service.(crawler.IteratorInterface)

			service, err = ctx.SafeGet("video_extractor")
			if err != nil {
				return nil, err
			}
			videoExtractor := service.(coreContract.VideoExtractorInterface)

			return commands.NewPrintCsvHandler(
				episodesPageIterator,
				pageItemsIterator,
				videoExtractor,
			).(cliContract.CommandHandlerInterface), nil
		},
	})

	return nil
}
