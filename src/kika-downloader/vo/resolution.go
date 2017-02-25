package vo

// Resolution resolution
type Resolution struct {
	width  int
	height int
}

// NewResolution return new Resolution value object
func NewResolution(x, y int) Resolution {
	return Resolution{x, y}
}

// GetWidth get width
func (r Resolution) GetWidth() int {
	return r.width
}

// GetHeight get height
func (r Resolution) GetHeight() int {
	return r.height
}
