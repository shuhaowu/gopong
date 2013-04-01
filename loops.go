package main

import (
	"fmt"
	"github.com/errnoh/wde.buffer"
	"github.com/skelterjohn/go.wde"
	"image"
	"image/draw"
	"time"
)

var game Game
var frames, updates uint

// Setup code.
func setup(window wde.Window, done chan bool) {
	fmt.Println("Setup")
	window.LockSize(true)
	window.Show()

	buffer.Create(window, image.Black)

	width, height := window.Size()
	game = InitializeGame(width, height)

	window.FlushImage()

	events := window.EventChan()

	go func() {
		for {
			time.Sleep(time.Second / 60)
			game.Update()
			updates++
		}
	}()

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

func update(window wde.Window, screen wde.Image) {
	for {
		r1 := game.Board1.Rect.Rectangle()
		r2 := game.Board2.Rect.Rectangle()

		buffer.Draw(r1, image.White, image.ZP, draw.Src)
		buffer.Draw(r2, image.White, image.ZP, draw.Src)
		buffer.Flip()
		frames++
	}
}

func cleanup() {
	fmt.Println("Cleanup")
}
