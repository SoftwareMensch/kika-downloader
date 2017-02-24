package http

import (
	"io/ioutil"
	"kika-downloader/config"
	"kika-downloader/http"
	testConfig "kika-downloader_test/config"
	"strings"
	"testing"
)

type ipInfoDTO struct {
	Ip string
}

func TestTorSocksProxy(t *testing.T) {
	appContext, err := config.SetupApp(testConfig.TorSocksProxyURL, "")
	if err != nil {
		t.Error(err)
	}

	service, err := appContext.SafeGet("http_client")
	if err != nil {
		t.Error(err)
	}
	client := service.(http.ClientInterface)

	resp, err := client.Get(testConfig.TorTestURL)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	jsResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	strResponse := string(jsResponse)
	expectedResponse := testConfig.TorTestValidResponse

	if !strings.Contains(strResponse, expectedResponse) {
		t.Error()
	}
}
