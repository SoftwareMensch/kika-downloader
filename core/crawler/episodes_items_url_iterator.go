package crawler

import (
	"fmt"
	"gopkg.in/xmlpath.v2"
	"log"
	"net/url"
	"rkl.io/kika-downloader/core/http"
)

// episodesItemsUrlIterator iterator for items of a page
type episodesItemsUrlIterator struct {
	abstractUrlIterator

	xpathEpisodesItems string
}

// NewEpisodesItemsUrlIterator returns iterator for episode items
func NewEpisodesItemsUrlIterator(
	client http.ClientInterface,
	xpathEpisodesItems string,
) IteratorInterface {
	i := &episodesItemsUrlIterator{
		xpathEpisodesItems: xpathEpisodesItems,
	}
	i.abstractUrlIterator.httpClient = client

	return i
}

// Run start iteration
func (i *episodesItemsUrlIterator) Run() <-chan string {
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

		path, err := xmlpath.Compile(i.xpathEpisodesItems)
		if err != nil {
			log.Println(err)
		}

		nodeIter := path.Iter(docRoot)

		for nodeIter.Next() {
			rawURL := nodeIter.Node().String()

			if _, err := url.Parse(rawURL); err != nil {
				// TODO, logging
				log.Printf("invalid url: %s\n", "fff")
				continue
			}

			i.urlChannel <- fmt.Sprintf(
				"%s://%s%s",
				i.crawlURL.Scheme,
				i.crawlURL.Host,
				rawURL,
			)
		}
	}()

	return i.urlChannel
}
