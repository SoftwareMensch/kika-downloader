package downloader

import (
	"fmt"
	"github.com/icrowley/fake"
	"kika-downloader/config"
	"kika-downloader/contract"
	"kika-downloader/model"
	"kika-downloader/vo"
	testConfig "kika-downloader_test/config"
	"net/url"
	"testing"
)

func TestCommonFileDownload(t *testing.T) {
	rawTestUrl := testConfig.DownloadTestURL
	testUrlTotalBytes := testConfig.DownloadTestURLTotalBytes

	appContext, err := config.InitApp(testConfig.TorSocksProxyURL, "")
	if err != nil {
		t.Error(err)
	}

	service, err := appContext.SafeGet("video_downloader")
	if err != nil {
		t.Error(err)
	}
	videoDownloader := service.(contract.VideoDownloaderInterface)

	parsedTestUrl, err := url.Parse(rawTestUrl)
	if err != nil {
		t.Error(err)
	}

	// create pseudo video entity
	dummyVideo := model.NewVideo(
		"12c70799-214a-4785-83cc-063289264fa4",
		fake.Word(),
		fake.Words(),
		fake.Sentence(),
		fake.Language(),
		fake.Day(),
		vo.NewResolution(fake.LatitudeMinutes(), fake.LatitudeDegress()),
		parsedTestUrl,
		testUrlTotalBytes,
	)

	fmt.Printf("rawTestUrl: %s\ntestUrlTotalBytes: %d\nvideoDownloader: %v\ndummyVideo: %v\n",
		rawTestUrl, testUrlTotalBytes, videoDownloader, dummyVideo)
}
