package contract

// VideoDownloaderInterface interface for downloader
type VideoDownloaderInterface interface {
	// Download start downloading a file
	Download(video VideoInterface, outputDir string) (<-chan IoProgressInterface, error)
}
