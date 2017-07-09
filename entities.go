package main

import (
	blt "github.com/dqnx/bearlibterminal"
	v2 "github.com/dqnx/vector2"
)

// Cell contains the necessary info for Blt to draw a cell in its grid.
type Cell struct {
	Position v2.Vector
	Color    uint32
	Code     int
}

// Drawable defines entities which can be drawn to the console.
type Drawable interface {
	// draw returns the color and code integers for use in blt.
	draw() Cell
}

// Moveable defines entities which have a changeable world position.
type Moveable interface {
	// Move will add a displacement vector to a position.
	move(delta v2.Vector)

	// Location returns the location of the movable in the game world.
	location() v2.Vector
}

// Concrete is a physical object in the game world, with position and graphics.
type Concrete struct {
	Position v2.Vector
	Code     int
	Color    uint32
}

// Tile is a container in the game world, representing the smallest unit of space.
type Tile struct {
	Concrete
	BlocksMovement bool
	BlocksSight    bool
	Dark           bool
	DarkCode       int
	DarkColor      uint32
}

func (t Tile) draw() Cell {
	c := Cell{Position: t.Position}
	if t.Dark {
		c.Color = t.DarkColor
		c.Code = t.DarkCode
	} else { // Light
		c.Color = t.Color
		c.Code = t.Code
	}
	return c
}

func newTile(tileType string, pos v2.Vector) Tile {
	argb := blt.ColorFromARGB

	t := Tile{Dark: false}
	t.Position = pos

	switch tileType {
	case "wall":
		t.Code = 0x0023                     // Hash
		t.DarkCode = 0x0023                 // Hash
		t.Color = argb(255, 232, 232, 232)  // Almost white
		t.DarkColor = argb(255, 75, 75, 75) // Dark grey
		t.BlocksMovement = true
		t.BlocksSight = true
	case "floor":
		t.Code = 0x00B7                     // Middle dot
		t.DarkCode = 0x0020                 // Blank space
		t.Color = argb(255, 232, 232, 232)  // Almost white
		t.DarkColor = argb(255, 15, 15, 15) // Almost black
		t.BlocksMovement = false
		t.BlocksSight = false
	}

	return t
}

// Actor holds values that describe a character, player, enemy, etc.
type Actor struct {
	Name string
	Concrete
}

func (a Actor) draw() Cell {
	return Cell{Position: a.Position, Color: a.Color, Code: a.Code}
}

func (a *Actor) move(delta v2.Vector) {
	a.Position = v2.Add(delta, a.Position)
}

func (a Actor) location() v2.Vector {
	return a.Position
}
