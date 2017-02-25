package extractor

import (
	"fmt"
	"kika-downloader/config"
	"kika-downloader/contract"
	testConfig "kika-downloader_test/config"
	"testing"
)

func TestVideoExtraction(t *testing.T) {
	appContext, err := config.SetupApp(testConfig.TorSocksProxyURL, "")
	if err != nil {
		t.Error(err)
	}

	service, err := appContext.SafeGet("video_extractor")
	if err != nil {
		t.Error(err)
	}

	metaDataExtractor := service.(contract.VideoExtractorInterface)

	video, err := metaDataExtractor.ExtractVideoFromURL(testConfig.ExtractorTestURL)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", video)
}
