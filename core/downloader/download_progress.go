package downloader

import (
	"fmt"
	"rkl.io/kika-downloader/core/contract"
)

type downloadProgress struct {
	totalByteCount   int64
	currentByteCount int64
}

// NewDownloadProgress return new download progress
func NewDownloadProgress(total, current int64) contract.IoProgressInterface {
	return &downloadProgress{total, current}
}

// GetTotalByteCount get total count of remote source file
func (p *downloadProgress) GetTotalByteCount() int64 {
	return p.totalByteCount
}

// GetCurrentByteCount get count of local bytes
func (p *downloadProgress) GetCurrentByteCount() int64 {
	return p.currentByteCount
}

// GetPercentage get current percentage
func (p *downloadProgress) GetPercentage() string {
	return fmt.Sprintf("%.2f", float32(p.currentByteCount)*100.0/float32(p.totalByteCount))
}
