package http

import (
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
)

// client http client
type client struct {
	http.Client

	socksProxyURL *url.URL
}

// NewClient return *client
func NewClient(socksProxyURL *url.URL) (ClientInterface, error) {
	client := &client{
		socksProxyURL: socksProxyURL,
	}

	if socksProxyURL != nil {
		dialer, err := proxy.FromURL(client.socksProxyURL, proxy.Direct)
		if err != nil {
			return nil, err
		}

		transport := &http.Transport{Dial: dialer.Dial}

		client.Transport = transport
	}

	return client, nil
}
