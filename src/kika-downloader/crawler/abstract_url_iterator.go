package crawler

import (
	"gopkg.in/xmlpath.v2"
	"kika-downloader/http"
	"net/url"
	"kika-downloader/utils"
)

type abstractUrlIterator struct {
	urlChannel chan string

	crawlURL   *url.URL
	httpClient http.ClientInterface
}

// GetUrlChannel return page channel
func (a *abstractUrlIterator) GetUrlChannel() <-chan string {
	return a.urlChannel
}

// SetCrawlingURL set url for crawling
func (a *abstractUrlIterator) SetCrawlingURL(crawlURL *url.URL) {
	a.crawlURL = crawlURL
}

func (a *abstractUrlIterator) domDocumentFromURL(u *url.URL) (*xmlpath.Node, error) {
	response, err := a.httpClient.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return utils.NodeFromHtmlBody(response.Body)
}
