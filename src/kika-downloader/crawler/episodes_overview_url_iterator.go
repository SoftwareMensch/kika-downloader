package crawler

import (
	"fmt"
	"gopkg.in/xmlpath.v2"
	"kika-downloader/http"
	"log"
	"net/url"
)

// episodesOverviewUrlIterator crawler for the episodes overview pages
type episodesOverviewUrlIterator struct {
	abstractUrlIterator

	xpathEpisodesOverviewPagePageItems string
}

// NewEpisodesOverviewUrlIterator return new *episodesOverviewUrlIterator
func NewEpisodesOverviewUrlIterator(
	client http.ClientInterface,
	xpathEpisodesOverviewPagePageItems string,

) IteratorInterface {
	i := &episodesOverviewUrlIterator{
		xpathEpisodesOverviewPagePageItems: xpathEpisodesOverviewPagePageItems,
	}
	i.abstractUrlIterator.httpClient = client

	return i
}

// Run start iteration
func (i *episodesOverviewUrlIterator) Run() <-chan string {
	i.urlChannel = make(chan string)

	go func() {
		defer close(i.urlChannel)

		// TODO, error logging/handling
		if i.crawlURL == nil {
			log.Println("no crawl url specified")
			return
		}

		docRoot, err := i.domDocumentFromURL(i.crawlURL)
		if err != nil {
			log.Println(err)
		}

		path, err := xmlpath.Compile(i.xpathEpisodesOverviewPagePageItems)
		if err != nil {
			log.Println(err)
		}

		uniqueUrls := []string{}

		sendUrlToPageChannel := func(rawURL string) {
			for _, u := range uniqueUrls {
				if u == rawURL {
					return
				}
			}

			uniqueUrls = append(uniqueUrls, rawURL)

			i.urlChannel <- fmt.Sprintf(
				"%s://%s%s",
				i.crawlURL.Scheme,
				i.crawlURL.Host,
				rawURL,
			)
		}

		nodeIter := path.Iter(docRoot)

		for nodeIter.Next() {
			rawURL := nodeIter.Node().String()

			if _, err := url.Parse(rawURL); err != nil {
				// TODO, logging
				log.Printf("invalid url: %s\n", rawURL)
				continue
			}

			sendUrlToPageChannel(rawURL)
		}
	}()

	return i.urlChannel
}
