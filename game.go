package main

import (
	"github.com/skelterjohn/go.wde"
	"github.com/shuhaowu/gopong/gt2d"
)

type Game struct {
	Board1 *Board
	Board2 *Board
	Width  int
	Height int
	Boundary *gt2d.Rectangle

	keysDown map[string]bool
}

func InitializeGame(width, height int) Game {
	var game Game
	game.Width = width
	game.Height = height

	game.Boundary = new(gt2d.Rectangle)
	game.Boundary.Min = gt2d.Vector2D{0, 0}
	game.Boundary.Max = gt2d.Vector2D{width+1, height+1} // correction factor.

	// Initializes the boards.
	// 50 is the height of the boards.
	game.Board1 = NewBoard(gt2d.Vector2D{80, height/2 - 50}, gt2d.Rect(0, 0, width / 2 + 1, height + 1))
	game.Board2 = NewBoard(gt2d.Vector2D{width - 90, height/2 - 50}, gt2d.Rect(width / 2, 0, width + 1, height + 1))

	game.keysDown = make(map[string]bool)
	game.keysDown[wde.KeyW] = false
	game.keysDown[wde.KeyS] = false
	game.keysDown[wde.KeyD] = false
	game.keysDown[wde.KeyA] = false
	game.keysDown[wde.KeyUpArrow] = false
	game.keysDown[wde.KeyDownArrow] = false
	game.keysDown[wde.KeyRightArrow] = false
	game.keysDown[wde.KeyLeftArrow] = false

	return game
}

func (game *Game) updateBoards() {
	var board1velocity, board2velocity gt2d.Vector2D

	if game.keysDown[wde.KeyW]{
		board1velocity.Y = -game.Board1.MaxVelocity
	} else if game.keysDown[wde.KeyS] {
		board1velocity.Y = game.Board1.MaxVelocity
	}

	if game.keysDown[wde.KeyD]{
		board1velocity.X = game.Board1.MaxVelocity
	} else if game.keysDown[wde.KeyA] {
		board1velocity.X = -game.Board1.MaxVelocity
	}

	if game.keysDown[wde.KeyUpArrow]{
		board2velocity.Y = -game.Board2.MaxVelocity
	} else if game.keysDown[wde.KeyDownArrow] {
		board2velocity.Y = game.Board2.MaxVelocity
	}

	if game.keysDown[wde.KeyLeftArrow] {
		board2velocity.X = -game.Board2.MaxVelocity
	} else if game.keysDown[wde.KeyRightArrow] {
		board2velocity.X = game.Board2.MaxVelocity
	}

	game.Board1.Move(board1velocity)
	game.Board2.Move(board2velocity)
}

func (game *Game) Update() {
	game.updateBoards()
}

func (game *Game) OnKeyDown(e wde.KeyDownEvent) {
	_, present := game.keysDown[e.Key]
	if present {
		game.keysDown[e.Key] = true
	}
}

func (game *Game) OnKeyUp(e wde.KeyUpEvent) {
	_, present := game.keysDown[e.Key]
	if present {
		game.keysDown[e.Key] = false
	}
}
