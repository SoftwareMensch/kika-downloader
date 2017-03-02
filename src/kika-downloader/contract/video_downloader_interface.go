package contract

// VideoDownloaderInterface interface for downloader
type VideoDownloaderInterface interface {
	// Download start downloading a file
	Download(video VideoInterface, outputDir string) (<-chan IoProgressInterface, error)

	// GetLocalFilePathAbs get absolute path of local downloaded file
	GetLocalFilePathAbs() string

	// GetLastError get last error
	GetLastError() error
}
