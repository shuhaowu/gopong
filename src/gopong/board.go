package main

import (
	"gt2d"
)

type Board struct {
	Rect        *gt2d.Rectangle
	Boundary    *gt2d.Rectangle
	MaxVelocity int
	Velocity    *gt2d.Vector2D
}

func NewBoard(bottomleft gt2d.Vector2D, boundary *gt2d.Rectangle) *Board {
	board := new(Board)
	board.Rect = new(gt2d.Rectangle)
	board.Rect.Min = bottomleft
	board.Rect.Max = gt2d.Vector2D{bottomleft.X + 15, bottomleft.Y + 100}
	board.Boundary = boundary
	board.MaxVelocity = 15
	return board
}

var ZV gt2d.Vector2D

func (board *Board) Move(v gt2d.Vector2D) {
	board.Rect.TranslateIPV(&v)
	if !board.Rect.In(board.Boundary) {
		board.Rect.TranslateIP(-v.X, -v.Y) // Problematic
		board.Velocity = &ZV
	} else {
		board.Velocity = &v
	}
}

// Since we don't really care about the top. left and right should do

const (
	LEFT  = "left"
	RIGHT = "right"
)

func (board *Board) Collide(rect *gt2d.Rectangle) string {
	if board.Rect.Collide(rect) {
		if rect.Max.X > board.Rect.Min.X && rect.Min.X <= board.Rect.Min.X {
			return LEFT
		} else {
			return RIGHT
		}
	}
	return ""
}
