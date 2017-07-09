package main

import v2 "github.com/dqnx/vector2"

type rect struct {
	X1, X2, Y1, Y2 int
}

// center returns the middle coordinates of a rect, rounding down fractions.
func (r rect) center() v2.Vector {
	var v v2.Vector
	v.X = int((r.X1 + r.X2) / 2)
	v.Y = int((r.Y1 + r.Y2) / 2)

	return v
}

func (r rect) intersect(r2 rect) bool {
	intersects := false
	if r.X1 <= r2.X2 && r.X2 >= r2.X1 && r.Y1 <= r2.Y2 && r.Y2 >= r2.Y1 {
		intersects = true
	}
	return intersects

}
func createRect(x, y, w, h int) rect {
	var r rect
	r.X1 = x
	r.Y1 = y
	r.X2 = x + w
	r.Y2 = y + h
	return r
}

type tileMap struct {
	Tiles [][]Tile
}

// createMap makes a 2D array of Tiles, representing the play area or map.
func createMap(size v2.Vector) tileMap {
	x, y := size.X, size.Y
	m := make([][]Tile, x)

	for i := 0; i < x; i++ {
		m[i] = make([]Tile, y)
		for j := 0; j < y; j++ {
			// Initialize all tiles to "wall"
			m[i][j] = newTile("wall", v2.Vector{X: i, Y: j})
		}
	}

	var t tileMap
	t.Tiles = m

	return t
}

// drawMap is a helper function to loop to each Tile and draw it.
func (t tileMap) draw(c chan Cell) {
	for _, outer := range t.Tiles {
		for _, inner := range outer {
			c <- inner.draw()
		}
	}
}

// get() returns a Tile pointer at the target vector.
func (t tileMap) get(v v2.Vector) *Tile {
	return &t.Tiles[v.X][v.Y]
}

// checkCollision determines if a movement will move to a blocked tile.
// In future, it should check the type of collision (unit, wall, item, etc.)
func (t tileMap) collision(m Moveable, delta v2.Vector) bool {
	newPos := v2.Add(m.location(), delta)
	target := t.get(newPos)
	return target.BlocksMovement
}

// createRoom changes a rectangle of space to floors, making a "room."
func (t tileMap) createRoom(r rect) {
	for x := r.X1; x <= r.X2; x++ {
		// Add 1 to the y start, to leave buffer wall between rooms.
		for y := r.Y1 + 1; y <= r.Y2; y++ {
			t.Tiles[x][y] = newTile("floor", v2.Vector{X: x, Y: y})
		}
	}
}

// min returns the minimum of 2 integers
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// max returns the maximum of 2 integers
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (t tileMap) tunnelHori(x1, x2, y int) {
	for x := min(x1, x2); x <= max(x1, x2); x++ {
		t.Tiles[x][y] = newTile("floor", v2.Vector{X: x, Y: y})
	}
}

func (t tileMap) tunnelVerti(y1, y2, x int) {
	for y := min(y1, y2); y <= max(y1, y2); y++ {
		t.Tiles[x][y] = newTile("floor", v2.Vector{X: x, Y: y})
	}
}
