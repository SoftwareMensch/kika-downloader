package crawler

import (
	"kika-downloader/config"
	"kika-downloader/crawler"
	testConfig "kika-downloader_test/config"
	"testing"
	"kika-downloader/http"
	"net/url"
	"fmt"
)

func TestPageIteration(t *testing.T) {
	appContext, err := config.SetupApp(testConfig.TorSocksProxyURL, testConfig.EpisodesOverviewURL)
	if err != nil {
		t.Error(err)
	}

	service, err := appContext.SafeGet("http_client")
	if err != nil {
		t.Error(err)
	}
	httpClient := service.(http.ClientInterface)

	service, err = appContext.SafeGet("episodes_overview_url_iterator")
	if err != nil {
		t.Error(err)
	}
	iterator := service.(crawler.IteratorInterface)

	gotValidURL := false

	// validate every url received from iterator
	for rawURL := range iterator.Run() {
		if _, err := url.Parse(rawURL); err != nil {
			t.Error(err)
		}

		if _, err := httpClient.Get(rawURL); err != nil {
			t.Error(err)
		}

		fmt.Printf("validated overview url: \"%s\"\n", rawURL)

		gotValidURL = true
	}

	if !gotValidURL {
		t.Error(fmt.Errorf("unable to retrieve a valid page url"))
	}
}
