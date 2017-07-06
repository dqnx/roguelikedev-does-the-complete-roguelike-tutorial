package main

import (
	v2 "github.com/dqnx/vector2"
)

type rect struct {
	X1, X2, Y1, Y2 int
}

func createRect(x, y, w, h int) rect {
	var r rect
	r.X1 = x
	r.Y1 = y
	r.X2 = x + w
	r.Y2 = y + h
	return r
}

func createRoom(t tileMap, r rect) {
	for x := r.X1; x <= r.X2; x++ {
		// Add 1 to the y start, to leave buffer wall between rooms.
		for y := r.Y1 + 1; y <= r.Y2; y++ {
			t.Tiles[x][y] = createTile("floor", v2.Vector{x, y})
		}
	}
}


