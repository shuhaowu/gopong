package main

import (
	"gt2d"
)

type Game struct {
	Rectangle gt2d.Rectangle
	Width int
	Height int
}

func InitializeGame(width, height int) Game {
	var game Game
	game.Rectangle = gt2d.Rectangle{gt2d.Vector2D{0, 0}, gt2d.Vector2D{15, 10}}
	game.Width = width
	game.Height = height
	return game
}

func (game *Game) Update() {
	game.Rectangle.TranslateIP(4, 4)
}