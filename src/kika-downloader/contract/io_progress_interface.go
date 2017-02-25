package contract

// IoProgressInterface interface for download progress
type IoProgressInterface interface {
	// GetSourceBytesCount get total count of remote source file
	GetSourceBytesCount() int

	// GetDestBytesCount get count of local bytes
	GetDestBytesCount() int
}
