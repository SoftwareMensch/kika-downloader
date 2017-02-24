package http

import "net/http"

// ClientInterface http client interface
type ClientInterface interface {
	Get(url string) (resp *http.Response, err error)
}
