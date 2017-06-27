// Package vector implements a simple 2 value container.
package vector2

// Vector holds 2 integer values: x and y.
type Vector struct {
	x, y int
}

// Subtracts each vector by its elements
func Sub(m, s Vector) Vector {
	return Vector{x: m.x-s.x, y: m.y-s.y}
}

// Adds each vector by its elements
func Add(vectors ...Vector) Vector {
	var sum Vector
	for _, vectors := range vectors {
		sum.x += vectors.x
		sum.y += vectors.y
	}
	return sum
}

// Returns the dot product of a vector
func Dot(a, b Vector) int {
	dotproduct := a.x*b.x + a.y*b.y
	return dotproduct
}
