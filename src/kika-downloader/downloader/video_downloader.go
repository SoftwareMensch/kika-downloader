package downloader

import (
	"fmt"
	"io"
	"kika-downloader/contract"
	"kika-downloader/http"
	"os"
	"path"
)

type videoDownloader struct {
	httpClient http.ClientInterface

	localFilePathAbs string
	lastError        error
}

type progressReader struct {
	io.Reader

	totalBytesCount   int64
	currentBytesCount int64

	downloadProgressChannel chan contract.IoProgressInterface
}

func (r *progressReader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	if err != nil {
		return n, err
	}

	r.currentBytesCount += int64(n)

	r.downloadProgressChannel <- NewDownloadProgress(r.totalBytesCount, r.currentBytesCount)

	return n, err
}

// NewVideoDownloader return new video downloader
func NewVideoDownloader(httpClient http.ClientInterface) contract.VideoDownloaderInterface {
	return &videoDownloader{httpClient, "", nil}
}

// GetLocalFilePathAbs get absolute path of local downloaded file
func (d *videoDownloader) GetLocalFilePathAbs() string {
	return d.localFilePathAbs
}

// GetLastError get last error
func (d *videoDownloader) GetLastError() error {
	return d.lastError
}

// Download start downloading a file
func (d *videoDownloader) Download(video contract.VideoInterface, outputDir string) (<-chan contract.IoProgressInterface, error) {
	if _, err := os.Stat(outputDir); err != nil {
		return nil, err
	}

	progressChan := make(chan contract.IoProgressInterface)

	// download goes here
	go func() {
		defer close(progressChan)

		resp, err := d.httpClient.Get(video.GetVideoOriginURL().String())
		if err != nil {
			d.lastError = err
			return
		}
		defer resp.Body.Close()

		outputFilePath := fmt.Sprintf("%s%s%s", outputDir, string(os.PathSeparator), d.filenameFromVideo(video))
		if _, err := os.Stat(outputFilePath); err == nil {
			d.lastError = fmt.Errorf("skipping file: %s", outputFilePath)
			return
		}

		output, err := os.Create(outputFilePath)
		if err != nil {
			d.lastError = err
			return
		}
		defer output.Close()

		bodyReader := &progressReader{
			resp.Body,
			video.GetFileSize(),
			0,
			progressChan,
		}

		writtenBytes, err := io.Copy(output, bodyReader)
		if err != nil {
			d.lastError = err
			return
		}

		progressChan <- NewDownloadProgress(video.GetFileSize(), writtenBytes)

		d.localFilePathAbs = outputFilePath
	}()

	return progressChan, nil
}

func (d *videoDownloader) filenameFromVideo(video contract.VideoInterface) string {
	originURL := video.GetVideoOriginURL()

	fileExt := path.Ext(originURL.String())

	return fmt.Sprintf("(%.2d) %s - %s%s",
		video.GetEpisodeNumber(),
		video.GetSeriesTitle(),
		video.GetEpisodeTitle(),
		fileExt,
	)
}
