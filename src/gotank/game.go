package main

import (
	"image"
	"fmt"
	"gt2d"
)

type Game struct {
	Rectangle image.Rectangle
	point gt2d.Vector2D
	Width int
	Height int
}

func InitializeGame(width, height int) Game {
	var game Game
	game.Rectangle = image.Rectangle{image.Point{0, 0}, image.Point{15, 10}}
	game.Width = width
	game.Height = height
	return game
}

func (game *Game) Update() {
	game.Rectangle = game.Rectangle.Add(image.Point{4, 4})
	fmt.Println(game.Rectangle)
}