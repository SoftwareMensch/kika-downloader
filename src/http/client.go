package http

import (
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
)

// Client http client
type Client struct {
	http.Client

	socksProxyURL *url.URL
}

// NewClient return *TorClient
func NewClient() (*Client, error) {
	parsedSocks5URL, err := url.Parse("socks5://192.168.242.1:9050")
	if err != nil {
		return nil, err
	}

	return &Client{
		socksProxyURL: parsedSocks5URL,
	}, nil
}

// ProxyGet get resource via proxy
func (c *Client) ProxyGet(rawURL string) (*http.Response, error) {
	dialer, err := proxy.FromURL(c.socksProxyURL, proxy.Direct)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{Dial: dialer.Dial}
	client := &http.Client{Transport: transport}

	return client.Get(rawURL)
}
