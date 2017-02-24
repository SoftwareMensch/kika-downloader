package services

import (
	"github.com/sarulabs/di"
	"kika-downloader/http"
	"net/url"
)

// AssignHttpClient assign http client service
func AssignHttpClient(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "http_client",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			torSocksProxyRawURL, err := ctx.SafeGet("tor_socks_proxy_url")
			if err != nil {
				return nil, err
			}

			torSocksProxyURL, err := url.Parse(torSocksProxyRawURL.(string))
			if err != nil {
				return nil, err
			}

			httpClient, err := http.NewClient(torSocksProxyURL)
			if err != nil {
				return nil, err
			}

			return httpClient.(http.ClientInterface), nil
		},
	})

	return nil
}
