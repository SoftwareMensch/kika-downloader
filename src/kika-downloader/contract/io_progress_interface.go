package contract

// IoProgressInterface interface for download progress
type IoProgressInterface interface {
	// GetTotalByteCount get total count of remote source file
	GetTotalByteCount() int64

	// GetCurrentByteCount get count of local bytes
	GetCurrentByteCount() int64

	// GetPercentage get current percentage
	GetPercentage() string
}
