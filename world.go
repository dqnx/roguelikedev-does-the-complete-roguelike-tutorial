package main

import (
	blt "github.com/dqnx/bearlibterminal"
	v2 "github.com/dqnx/vector2"
)

type tileMap struct {
	Tiles [][]Tile
}

// createMap makes a 2D array of Tiles, representing the play area or map.
func createMap(size v2.Vector) tileMap {
	x, y := size.x, size.y
	m := make([][]Tile, x)
	
	for i := 0; i < x; i++ {
		m[i] = make([]int, y)
		for j := 0; j < y; j++ {
			// Initialize all tiles to "wall"
			m[i][j] = createTile("wall", v2.Vector{i, j})
		}
	}
	
	var t tileMap
	t.Tiles = m
	
	return t
}

// drawMap is a helper function to loop to each Tile and draw it.
func (t tileMap) draw() {
	for _, outer := range t {
		for _,  inner := range outer {
			color, code := inner.draw()
			blt.Color(blt.ColorFromName(color))
			blt.Put(inner.Position.X, inner.Position.Y, code)
		}
	}
}

// get() returns a Tile pointer at the target vector.
func (t tileMap) get(v v2.Vector) *Tile {
	return &t.Tiles[v.X][x.Y]
}

// checkCollision determines if a movement will move to a blocked tile.
// In future, it should check the type of collision (unit, wall, item, etc.)
func (t tileMap) collision(m Moveable, delta v2.Vector) bool {
	newPos := v2.Add(m.Location(), delta)
	target := t.get(newPos)
	
	return target.BlocksMovement
}
