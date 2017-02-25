package downloader

import (
	"kika-downloader/contract"
	"kika-downloader/http"
	"os"
)

type videoDownloader struct {
	httpClient http.ClientInterface
}

// NewVideoDownloader return new video downloader
func NewVideoDownloader(httpClient http.ClientInterface) contract.VideoDownloaderInterface {
	return &videoDownloader{httpClient}
}

// Download start downloading a file
func (d *videoDownloader) Download(video contract.VideoInterface, outputDir string) (<-chan contract.IoProgressInterface, error) {
	if _, err := os.Stat(outputDir); err != nil {
		return nil, err
	}

	return nil, nil
}
