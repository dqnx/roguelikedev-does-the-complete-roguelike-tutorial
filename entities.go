package main

import (
	v2 "github.com/dqnx/vector2"
)

// Drawable defines entities which can be drawn to the console.
type Drawable interface {
	draw() int
}

// Movable defines entities which have a changeable world position.
type Movable interface {
	// Move will add a displacement vector to a position.
	move(delta v2.Vector)
	
	// Location returns the location of the movable in the game world.
	location() v2.Vector
}

// Concrete is a physical object in the game world, with position and graphics.
type Concrete struct {
	Position v2.Vector
	Code int
	Color string
}

// Actor holds values that describe a character, player, enemy, etc.
type Actor struct {
	Position v2.Vector
	Concrete
}

func (a Actor) draw() string, int {
	return a.Color, a.Code
}

func (a *Actor) move(delta v2.Vector) {
	a.Position = v2.Add(delta, a.Position)
}
