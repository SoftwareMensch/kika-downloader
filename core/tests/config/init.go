package config

import (
	"github.com/sarulabs/di"
	"rkl.io/kika-downloader/core/config"
)

// InitTestContext basic app set up
func InitTestContext(socksProxyUrl string) (di.Context, error) {
	builder, err := config.InitCoreBuilder(socksProxyUrl)
	if err != nil {
		return nil, err
	}

	return builder.Build(), nil
}
