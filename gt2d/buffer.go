package gt2d

import (
	"image"
)

type ScreenBuffer struct {
	*image.RGBA
	Background *image.RGBA
}

func NewScreenBuffer(width, height int, background *image.RGBA) *ScreenBuffer {
	buffer := new(ScreenBuffer)
	buffer.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))
	buffer.Background = background
	return buffer
}

func (buffer *ScreenBuffer) Reset() {
	copy(buffer.Pix, buffer.Background.Pix)
}

// This is here because we want to eventually make this BGRA