package main

import (
	v2 "github.com/dqnx/vector2"
)

// Drawable defines entities which can be drawn to the console.
type drawable interface {
	draw() int
}

// Movable defines entities which have a changeable world position.
type movable interface {
	move(dx, dy int)
}

// Actor holds values that describe a character, player, enemy, etc.
type actor struct {
	Position v2.Vector
	Name     string
	Code     int
}

func (a actor) draw() int {
	return a.Code
}

func (a *actor) move(dx, dy int) {
	a.Position.X += dx
	a.Position.Y += dy
}
