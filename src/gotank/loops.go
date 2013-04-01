package main

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	"image/color"
)

var game Game

func resetScreen(screen wde.Image, width int, height int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			screen.Set(x, y, color.Black)
		}
	}
}

// Setup code.
func setup(window wde.Window, done chan bool) {
	fmt.Println("Setup")
	window.LockSize(true)
	window.Show()

	width, height := window.Size()
	screen := window.Screen()

	resetScreen(screen, width, height)
	window.FlushImage()

	events := window.EventChan()
	// Events
	go func() {
	loop:
		for ei := range events {
			switch ei.(type) {
			case wde.CloseEvent:
				break loop
			}
		}

		done <- true
	}()

	game = InitializeGame(width, height)
}



func update(window wde.Window, screen wde.Image) bool {
	resetScreen(screen, game.Width, game.Height)
	game.Update()
	DrawRectangle(screen, &(game.Rectangle), color.White)
	return true
}

func cleanup() {
	fmt.Println("Cleanup")
}
