package main

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	"image/color"
	"gt2d"
	"image"
)

var game Game
var buffer *gt2d.ScreenBuffer

// Setup code.
func setup(window wde.Window, done chan bool) {
	fmt.Println("Setup")
	window.LockSize(true)
	window.Show()

	width, height := window.Size()
	game = InitializeGame(width, height)
	background := image.NewRGBA(image.Rect(0, 0, width, height))
	buffer = gt2d.NewScreenBuffer(width, height, background)

	screen := window.Screen()

	screen.CopyRGBA(buffer.RGBA, buffer.Rect)
	window.FlushImage()

	events := window.EventChan()
	// Events
	go func() {
	loop:
		for ei := range events {
			switch e := ei.(type) {
			case wde.CloseEvent:
				break loop
			case wde.KeyDownEvent:
				game.OnKeyDown(e)
			case wde.KeyUpEvent:
				game.OnKeyUp(e)
			}
		}

		done <- true
	}()
}

func update(window wde.Window, screen wde.Image) bool {
	buffer.Reset()
	game.Update()
	gt2d.DrawRectangle(buffer, game.Board1.Rect, color.White)
	gt2d.DrawRectangle(buffer, game.Board2.Rect, color.White)
	screen.SetRGBA(buffer.RGBA)
	return true
}

func cleanup() {
	fmt.Println("Cleanup")
}
