package util

import "strconv"

func Sliceatoi(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, v := range strings {
		var err error
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}
