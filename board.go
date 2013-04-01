package main

import (
	"github.com/shuhaowu/gopong/gt2d"
)

type Board struct {
	Rect *gt2d.Rectangle
	Boundary *gt2d.Rectangle
	MaxVelocity int
}

func NewBoard(bottomleft gt2d.Vector2D, boundary *gt2d.Rectangle) *Board {
	board := new(Board)
	board.Rect = new(gt2d.Rectangle)
	board.Rect.Min = bottomleft
	board.Rect.Max = gt2d.Vector2D{bottomleft.X + 10, bottomleft.Y + 100}
	board.Boundary = boundary
	board.MaxVelocity = 10
	return board
}

func (board *Board) Move(v gt2d.Vector2D) {
	board.Rect.TranslateIPV(&v)
	if !board.Rect.In(board.Boundary){
		board.Rect.TranslateIP(-v.X, -v.Y)
	}
}
