package main

import (
	"gt2d"
	"math/rand"
	"time"
)

type Bullet struct {
	Rect            *gt2d.Rectangle
	Boundary        *gt2d.Rectangle
	CurrentPosition *gt2d.Vector2D
	LastPosition    *gt2d.Vector2D
	Velocity        *gt2d.Vector2D
}

func NewBullet(boundary *gt2d.Rectangle) *Bullet {
	rand.Seed(time.Now().UTC().UnixNano())

	bullet := new(Bullet)
	bullet.Boundary = boundary

	// init with a random velocity and direction
	bullet.Velocity = new(gt2d.Vector2D)

	for {
		if abs(bullet.Velocity.X) > 1 && abs(bullet.Velocity.Y) > 1 {
			break
		}
		bullet.Velocity.X = rand.Intn(16) - 8
		bullet.Velocity.Y = rand.Intn(16) - 8
	}

	bullet.CurrentPosition = new(gt2d.Vector2D)
	bullet.LastPosition = new(gt2d.Vector2D)
	bullet.CurrentPosition.X = boundary.Width() / 2
	bullet.CurrentPosition.Y = boundary.Height() / 2
	bullet.LastPosition.X = bullet.CurrentPosition.X
	bullet.LastPosition.Y = bullet.CurrentPosition.Y

	bullet.Rect = gt2d.Rect(bullet.CurrentPosition.X-5, bullet.CurrentPosition.Y-5, bullet.CurrentPosition.X+5, bullet.CurrentPosition.Y+5)
	return bullet
}

func (bullet *Bullet) Update() {
	bullet.LastPosition.X = bullet.CurrentPosition.X
	bullet.LastPosition.Y = bullet.CurrentPosition.Y
	bullet.CurrentPosition.AddIP(bullet.Velocity)
	bullet.Rect.TranslateIPV(bullet.Velocity)
}
