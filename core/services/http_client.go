package services

import (
	"github.com/sarulabs/di"
	"net/url"
	"rkl.io/kika-downloader/core/http"
)

// AssignHttpClient assign http client service
func AssignHttpClient(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "http_client",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			var socksProxyURL *url.URL

			socksProxyRawURL, err := ctx.SafeGet("socks_proxy_url")
			if err != nil {
				return nil, err
			}

			if socksProxyRawURL.(string) != "" {
				socksProxyURL, err = url.Parse(socksProxyRawURL.(string))
				if err != nil {
					return nil, err
				}
			}

			httpClient, err := http.NewClient(socksProxyURL)
			if err != nil {
				return nil, err
			}

			return httpClient.(http.ClientInterface), nil
		},
	})

	return nil
}
