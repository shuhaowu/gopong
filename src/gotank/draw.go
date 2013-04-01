package main

import (
	"image"
	"image/color"
	"github.com/skelterjohn/go.wde"
)

func DrawRectangle(screen wde.Image, rectangle *image.Rectangle, color color.Color) {
	for x := (*rectangle).Min.X; x < (*rectangle).Max.X; x++ {
		for y := (*rectangle).Min.Y; y < (*rectangle).Max.Y; y++ {
			screen.Set(x, y, color)
		}
	}
}