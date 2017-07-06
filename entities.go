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
	Color int
}

// Tile is a container in the game world, representing the smallest unit of space.
type Tile struct {
	Concrete
	BlocksMovement bool
	BlocksSight bool
	Dark bool
	DarkCode int
	DarkColor string
}

func (t Tile) draw() string, int {
	if dark {
		return t.DarkColor, t.DarkCode	
	}
	return t.Color, t.Code
}

func createTile(tileType string, pos v2.Vector) Tile {
	t := Tile{X: pos.X, Y: pos.Y, Dark: false}
	switch tileType {
		case "wall":
			t.Code = 0x0023				// Hash
			t.DarkCode = 0x0023			// Hash
			t.Color = argb(255, 232, 232, 232) 	// Almost white
			t.DarkColor = argb(255, 75, 75, 75)	// Dark grey
			t.BlocksMovement = true
			t.BlocksSight = true
		case "floor":
			t.Code = 0x00B7 			// Middle dot
			t.DarkCode = 0x0020 			// Blank space
			t.Color = argb(255, 232, 232, 232) 	// Almost white
			t.DarkColor = argb(255, 15, 15, 15)	// Almost black
			t.BlocksMovement = false
			t.BlocksSight = false
	}
}

// Actor holds values that describe a character, player, enemy, etc.
type Actor struct {
	Name string
	Concrete
}

func (a Actor) draw() string, int {
	return a.Color, a.Code
}

func (a *Actor) move(delta v2.Vector) {
	a.Position = v2.Add(delta, a.Position)
}
