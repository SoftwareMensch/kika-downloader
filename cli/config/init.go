package config

import (
	"github.com/sarulabs/di"
	cliServices "rkl.io/kika-downloader/cli/services"
	"rkl.io/kika-downloader/core/config"
)

// InitCliContext basic app set up
func InitCliContext(socksProxyUrl string) (di.Context, error) {
	builder, err := config.InitCoreBuilder(socksProxyUrl)
	if err != nil {
		return nil, err
	}

	// fetch all command handler
	if err := cliServices.AssignFetchAllCommandHandler(builder); err != nil {
		return nil, err
	}

	// print csv command handler
	if err := cliServices.AssignPrintCsvCommandHandler(builder); err != nil {
		return nil, err
	}

	return builder.Build(), nil
}
