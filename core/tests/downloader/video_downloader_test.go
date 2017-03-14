package downloader

import (
	"encoding/hex"
	"fmt"
	"github.com/icrowley/fake"
	"io/ioutil"
	"net/url"
	"os"
	"rkl.io/kika-downloader/config"
	"rkl.io/kika-downloader/core/contract"
	"rkl.io/kika-downloader/core/model"
	testConfig "rkl.io/kika-downloader/core/tests/config"
	"rkl.io/kika-downloader/core/utils"
	"rkl.io/kika-downloader/core/vo"
	"testing"
)

func TestVideoFileDownload(t *testing.T) {
	rawTestUrl := testConfig.DownloadTestURL
	testUrlTotalBytes := int64(testConfig.DownloadTestURLTotalBytes)

	appContext, err := config.InitApp("")
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

	outputTmpDir, err := ioutil.TempDir("", "vd-test-")
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

	localFile, err := download(videoDownloader, dummyVideo, outputTmpDir)
	if err != nil {
		t.Error(err)
	}

	if err = videoDownloader.GetLastError(); err != nil {
		t.Error(err)
	}

	// test if already downloaded files will skipped
	_, err = download(videoDownloader, dummyVideo, outputTmpDir)
	if videoDownloader.GetLastError() == nil {
		t.Errorf("already downloaded files shoud be skipped")
	}

	// cleanup downloaded file
	if err = os.Remove(localFile); err != nil {
		t.Error(err)
	}

	// cleanup temp directory
	if err = os.Remove(outputTmpDir); err != nil {
		t.Error(err)
	}
}

func download(videoDownloader contract.VideoDownloaderInterface, dummyVideo contract.VideoInterface, outputTmpDir string) (string, error) {
	progressChannel, err := videoDownloader.Download(dummyVideo, outputTmpDir)
	if err != nil {
		return "", err
	}

	// show progress
	for progress := range progressChannel {
		if err := videoDownloader.GetLastError(); err != nil {
			return "", err
		}

		fmt.Printf("progress: %s %%\n", progress.GetPercentage())
	}

	// test if file was really downloaded
	fileInfo, err := os.Stat(videoDownloader.GetLocalFilePathAbs())
	if err != nil {
		return "", err
	}

	// test if file has expected size
	if fileInfo.Size() != testConfig.DownloadTestURLTotalBytes {
		return "", fmt.Errorf("downloaded file size does not match expected one")
	}

	// test if expected checksum matches
	sha256Sum, err := utils.SHA256FromFile(videoDownloader.GetLocalFilePathAbs())
	if err != nil {
		return "", err
	}

	sha256SumString := hex.EncodeToString(sha256Sum)
	expectedSha256SumString := testConfig.DownloadTestUrlSha256Sum

	if sha256SumString != expectedSha256SumString {
		return "", fmt.Errorf("wrong downloaded file checksum")
	}

	return videoDownloader.GetLocalFilePathAbs(), nil
}
