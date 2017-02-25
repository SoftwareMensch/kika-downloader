package downloader

import "kika-downloader/contract"

type downloadProgress struct {
	totalByteCount   int
	currentByteCount int
}

// NewDownloadProgress return new download progress
func NewDownloadProgress(total, current int) contract.IoProgressInterface {
	return &downloadProgress{total, current}
}

// GetSourceBytesCount get total count of remote source file
func (p *downloadProgress) GetSourceBytesCount() int {
	return p.totalByteCount
}

// GetDestBytesCount get count of local bytes
func (p *downloadProgress) GetDestBytesCount() int {
	return p.currentByteCount
}
