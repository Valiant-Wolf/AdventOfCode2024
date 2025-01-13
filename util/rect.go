package util

import "aoc24/vect"

// A Rect represents a rectangle in 2D space.
// Each dimension is inclusive-exclusive.
type Rect struct {
	left, top, right, bottom int
}

func NewRect(a, b vect.Vector) Rect {
	left, right := min(a.X, b.X), max(a.X, b.X)
	top, bottom := min(a.Y, b.Y), max(a.Y, b.Y)
	return Rect{left, top, right, bottom}
}

func (r Rect) Contains(v vect.Vector) bool {
	return v.X >= r.left && v.X < r.right && v.Y >= r.top && v.Y < r.bottom
}
