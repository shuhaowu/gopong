package main

import (
	"github.com/skelterjohn/go.wde"
	"gt2d"
	"fmt"
)


func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

type Game struct {
	Board1   *Board
	Board2   *Board
	Bul      *Bullet
	Width    int
	Height   int
	Boundary *gt2d.Rectangle
	currentFrame int
	lastCollision int
	Running bool

	keysDown map[string]bool
}

func InitializeGame(width, height int) Game {
	var game Game
	game.Width = width
	game.Height = height

	game.Boundary = new(gt2d.Rectangle)
	game.Boundary.Min = gt2d.Vector2D{0, 0}
	game.Boundary.Max = gt2d.Vector2D{width + 1, height + 1} // correction factor.

	// Initializes the boards.
	// 50 is the height of the boards.
	game.Board1 = NewBoard(gt2d.Vector2D{80, height/2 - 50}, gt2d.Rect(0, 0, width/2+1, height+1))
	game.Board2 = NewBoard(gt2d.Vector2D{width - 90, height/2 - 50}, gt2d.Rect(width/2+21, 0, width+1, height+1))
	game.Bul = NewBullet(game.Boundary)

	game.keysDown = make(map[string]bool)
	game.keysDown[wde.KeyW] = false
	game.keysDown[wde.KeyS] = false
	game.keysDown[wde.KeyD] = false
	game.keysDown[wde.KeyA] = false
	game.keysDown[wde.KeyUpArrow] = false
	game.keysDown[wde.KeyDownArrow] = false
	game.keysDown[wde.KeyRightArrow] = false
	game.keysDown[wde.KeyLeftArrow] = false

	game.Running = true

	return game
}

func (game *Game) updateBoards() {
	var board1velocity, board2velocity gt2d.Vector2D

	if game.keysDown[wde.KeyW] {
		board1velocity.Y = -game.Board1.MaxVelocity
	} else if game.keysDown[wde.KeyS] {
		board1velocity.Y = game.Board1.MaxVelocity
	}

	if game.keysDown[wde.KeyD] {
		board1velocity.X = game.Board1.MaxVelocity
	} else if game.keysDown[wde.KeyA] {
		board1velocity.X = -game.Board1.MaxVelocity
	}

	if game.keysDown[wde.KeyUpArrow] {
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

func (game *Game) updateBullet() {
	game.Bul.Update()
}


const MAX_BULLET_V = 12

// For our shitty simulator, v1 won't change. v2 gets reflected and some extra
// speed added. The MAX V is 15 15.
func collideAndChangeVelocity (v1 *gt2d.Vector2D, v2 *gt2d.Vector2D) {
	x := abs(v2.X) + abs(v1.X) / 4
	y := v2.Y + v1.Y / 2

	if x > MAX_BULLET_V {
		x = MAX_BULLET_V
	}

	// so bad. We need a limit functionality.
	if y > MAX_BULLET_V {
		y = 15
	} else if y < -MAX_BULLET_V {
		y = -MAX_BULLET_V
	}

	if v2.X < 0 {
		v2.X = x
	} else {
		v2.X = -x
	}

	v2.Y = y
}

// This method will set the bullet's speed
func (game *Game) detectCollision() {
	if game.Board1.Rect.Collide(game.Bul.Rect) && game.currentFrame - game.lastCollision > 20 {
		collideAndChangeVelocity(game.Board1.Velocity, game.Bul.Velocity)
		game.lastCollision = game.currentFrame
	}

	if game.Board2.Rect.Collide(game.Bul.Rect) && game.currentFrame - game.lastCollision > 20 {
		collideAndChangeVelocity(game.Board2.Velocity, game.Bul.Velocity)
		game.lastCollision = game.currentFrame
	}

	if game.Bul.Rect.Max.Y > game.Boundary.Max.Y || game.Bul.Rect.Min.Y < game.Boundary.Min.Y {
		game.Bul.Velocity.Y = -game.Bul.Velocity.Y
	}

	if game.Bul.Rect.Min.X > game.Boundary.Max.X {
		game.Running = false
		fmt.Println("Winner is the player on the left!")
	}

	if game.Bul.Rect.Max.X < game.Boundary.Min.X {
		game.Running = false
		fmt.Println("Winner is the player on the right!")
	}
}

func (game *Game) Update() {
	if game.Running {
		game.updateBoards()
		game.updateBullet()
		game.detectCollision()
		game.currentFrame++
	}
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
