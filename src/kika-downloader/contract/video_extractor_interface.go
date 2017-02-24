package contract

// VideoExtractorInterface interface for all video extractors
type VideoExtractorInterface interface {
	ExtractVideoFromURL(rawURL string) (VideoInterface, error)
}