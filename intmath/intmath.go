package intmath

import "strconv"

func Abs(val int) int {
	return max(val, -val)
}

func Sign(val int) int {
	switch {
	case val == 0:
		return 0
	case val < 0:
		return -1
	default:
		return 1
	}
}

func Digits(val int) int {
	return len(strconv.Itoa(Abs(val)))
}
