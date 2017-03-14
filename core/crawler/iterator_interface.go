package crawler

import "net/url"

type IteratorInterface interface {
	Run() <-chan string
	SetCrawlingURL(crawlURL *url.URL)
}
