package downloader

import (
	"encoding/hex"
	"fmt"
	"github.com/icrowley/fake"
	"io/ioutil"
	"kika-downloader/config"
	"kika-downloader/contract"
	"kika-downloader/model"
	"kika-downloader/utils"
	"kika-downloader/vo"
	testConfig "kika-downloader_test/config"
	"net/url"
	"os"
	"testing"
)

func TestVideoFileDownload(t *testing.T) {
	rawTestUrl := testConfig.DownloadTestURL
	testUrlTotalBytes := int64(testConfig.DownloadTestURLTotalBytes)

	appContext, err := config.InitApp("", "")
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

	outputTmpDir, err := ioutil.TempDir("", "vd-test-")
	if err != nil {
		t.Error(err)
	}

	progressChannel, err := videoDownloader.Download(dummyVideo, outputTmpDir)
	if err != nil {
		t.Error(err)
	}

	// show progress
	for progress := range progressChannel {
		if err := videoDownloader.GetLastError(); err != nil {
			t.Error(err)
		}

		fmt.Printf("progress: %s %%\n", progress.GetPercentage())
	}

	// test if file was really downloaded
	fileInfo, err := os.Stat(videoDownloader.GetLocalFilePathAbs())
	if err != nil {
		t.Error(err)
	}

	// test if file has expected size
	if fileInfo.Size() != testConfig.DownloadTestURLTotalBytes {
		t.Errorf("downloaded file size does not match expected one")
	}

	// test if expected checksum matches
	sha256Sum, err := utils.SHA256FromFile(videoDownloader.GetLocalFilePathAbs())
	if err != nil {
		t.Error(err)
	}

	sha256SumString := hex.EncodeToString(sha256Sum)
	expectedSha256SumString := testConfig.DownloadTestUrlSha256Sum

	if sha256SumString != expectedSha256SumString {
		t.Errorf("wrong downloaded file checksum")
	}

	// cleanup downloaded file
	if err = os.Remove(videoDownloader.GetLocalFilePathAbs()); err != nil {
		t.Error(err)
	}

	// cleanup temp directory
	if err = os.Remove(outputTmpDir); err != nil {
		t.Error(err)
	}
}
