package main

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"time"
)

func main() {
	go runGame()
	wde.Run()
}

func runGame() {
	width := 800
	height := 600

	window, err := wde.NewWindow(width, height)
	if err != nil {
		fmt.Println(err)
		return
	}
	window.SetTitle("Go Pong") // could probably be extended.
	window.SetSize(width, height)

	done := make(chan bool)
	setup(window, done)

	var totalFrames uint

	go update(window, window.Screen())

loop:
	for {
		select {
		case <-done:
			break loop
		case <-time.After(time.Second):
			fmt.Printf("FPS: %d UPS %d\n", frames, updates)
			totalFrames += frames
			frames = 0
			updates = 0
		}
	}

	cleanup()
	window.Close()
	fmt.Println("Update loop ran", totalFrames, "times")

	wde.Stop()
}
