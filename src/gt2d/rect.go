package gt2d

import (
	"image/color"
)

type Rectangle struct {
	Min, Max Vector2D
}

var NullRect Rectangle

// TODO: Add a way to fix the points.

func Rect(x1, y1, x2, y2 int) *Rectangle{
	return &(Rectangle{Vector2D{x1, y1}, Vector2D{x2, y2}})
}

func (rectangle *Rectangle) Translate(x, y int) Rectangle {
	return rectangle.TranslateV(&Vector2D{x, y})
}

func (rectangle *Rectangle) TranslateIP(x, y int) {
	rectangle.Min.X += x
	rectangle.Max.X += x
	rectangle.Min.Y += y
	rectangle.Max.Y += y
}

func (rectangle *Rectangle) TranslateV(vector *Vector2D) Rectangle {
	return Rectangle{rectangle.Min.Add(vector), rectangle.Max.Add(vector)}
}

func (rectangle *Rectangle) TranslateIPV(vector *Vector2D) {
	rectangle.Min.AddIP(vector)
	rectangle.Max.AddIP(vector)
}

func (rectangle *Rectangle) Width() int {
	return abs(rectangle.Max.X - rectangle.Min.X)
}

func (rectangle *Rectangle) Height() int {
	return abs(rectangle.Max.Y - rectangle.Min.Y)
}

func (rectangle *Rectangle) In(o *Rectangle) bool {
	return o.Min.X <= rectangle.Min.X && rectangle.Max.X <= o.Max.X && o.Min.Y <= rectangle.Min.Y && rectangle.Max.Y <= o.Max.Y
}

// Intersect returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (rectangle Rectangle) Intersect(s *Rectangle) Rectangle {
	// rectangle is already a copy.
	if rectangle.Min.X < s.Min.X {
	rectangle.Min.X = s.Min.X
	}
	if rectangle.Min.Y < s.Min.Y {
		rectangle.Min.Y = s.Min.Y
	}
	if rectangle.Max.X > s.Max.X {
		rectangle.Max.X = s.Max.X
	}
	if rectangle.Max.Y > s.Max.Y {
		rectangle.Max.Y = s.Max.Y
	}
	if rectangle.Min.X > rectangle.Max.X || rectangle.Min.Y > rectangle.Max.Y {
		return NullRect
	}
	return rectangle
}

// Union returns the smallest rectangle that contains both r and s.
func (r Rectangle) Union(s Rectangle) Rectangle {
	if r.Min.X > s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y > s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X < s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y < s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	return r
}

// Now is the section for drawing these rectangles to an image.

// Draws a rectangle onto an image. img will be modified.
// The rectangle will be clipped if it lands outside the image.
func DrawRectangle(buffer *ScreenBuffer, rectangle *Rectangle, color color.Color) {
	for x := rectangle.Min.X; x < rectangle.Max.X; x++ {
		for y := rectangle.Min.Y; y < rectangle.Max.Y; y++ {
			buffer.Set(x, y, color)
		}
	}
}