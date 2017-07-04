package main

import (
	v2 "github.com/dqnx/vector2"
)

type drawable interface {
	draw() int
}

type actor struct {
	Position v2.Vector
	Name     string
	Code     int
}

func (a actor) draw() int {
	return a.Code
}
