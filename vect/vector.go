package vect

import "fmt"

// A Vector represents a pair of integer dimensions.
type Vector struct {
	X int
	Y int
}

func (v Vector) Add(addend Vector) Vector {
	return Vector{v.X + addend.X, v.Y + addend.Y}
}

func (v Vector) Sub(subtrahend Vector) Vector {
	return Vector{v.X - subtrahend.X, v.Y - subtrahend.Y}
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

func Up() Vector {
	return Vector{0, -1}
}

func Down() Vector {
	return Vector{0, 1}
}

func Left() Vector {
	return Vector{-1, 0}
}

func Right() Vector {
	return Vector{1, 0}
}
