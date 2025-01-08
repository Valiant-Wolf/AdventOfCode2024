package util

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
