package main

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"sync"
	"time"
)

func main() {
	go runGame()
	wde.Run()
}

func runGame() {
	var wait sync.WaitGroup
	width := 800
	height := 600

	game := func() {
		window, err := wde.NewWindow(width, height)
		if err != nil {
			fmt.Println(err)
			return
		}
		window.SetTitle("Go Tank Game") // could probably be extended.
		window.SetSize(width, height)

		done := make(chan bool)
		setup(window, done)

		var i uint
		var currentTime int64
		var lastTime int64
		var totalFrames uint

		lastTime = time.Now().UnixNano()
	loop:
		for i = 0; ; i++ {
			currentTime = time.Now().UnixNano()
			if currentTime-lastTime > 1000000000 {
				lastTime = currentTime
				totalFrames += i
				fmt.Println("FPS:", i)
				i = 0
			}

			if !update(window, window.Screen()) {
				break loop
			}

			window.FlushImage()

			select {
			case <-done:
				break loop
			case <-time.After(time.Duration(20 * time.Millisecond)):
				continue
			}
		}

		cleanup()
		window.Close()
		fmt.Println("Update loop ran", totalFrames, "times")
		wait.Done()
	}

	wait.Add(1)
	go game()
	wait.Wait()
	wde.Stop()
}
