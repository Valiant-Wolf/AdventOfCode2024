package _1_historian_hysteria

import (
	"aoc24/intmath"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func readLists() (left []int, right []int) {
	// for each line in the input
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		// split the line into fields (separated by whitespace)
		fields := strings.Fields(line)

		// convert the left field to an int and append it to the left slice
		l, _ := strconv.Atoi(fields[0])
		left = append(left, l)

		// likewise for the right field
		r, _ := strconv.Atoi(fields[1])
		right = append(right, r)
	}

	return
}

func A() error {
	// read input into slices
	left, right := readLists()

	// sort the slices
	slices.Sort(left)
	slices.Sort(right)

	var sumDiff int

	// for each index in the left slice
	for i := range left {
		// calculate the difference between entries and add it to the total
		sumDiff += intmath.Abs(left[i] - right[i])
	}

	fmt.Println(sumDiff)

	return nil
}

func B() error {
	// read input into slices
	left, right := readLists()

	counts := make(map[int]int)

	// for each entry in the right slice
	for _, val := range right {
		// add to the count for that value
		counts[val] = counts[val] + 1
	}

	var similarity int

	// for each entry in the left slice
	for _, val := range left {
		// add to similarity the value times the count of that value in the right slice
		similarity += val * counts[val]
	}

	fmt.Println(similarity)

	return nil
}
