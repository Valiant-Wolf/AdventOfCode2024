package _4_ceres_search

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func readGrid() (grid [][]rune) {
	// for each line in the input
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		// append the runes of the line to the grid
		grid = append(grid, []rune(strings.TrimSpace(line)))
	}

	return
}

func kernel(input [][]rune, size int) <-chan [][]rune {
	// calculate the dimensions of the input padded so no kernels are out of range
	padding := size - 1
	paddedHeight := len(input) + padding*2
	paddedWidth := len(input[0]) + padding*2

	// initialise a padded source slice
	source := make([][]rune, paddedHeight)
	for row := range source {
		source[row] = make([]rune, paddedWidth)
	}

	// copy the input into the centre of the source slice
	for row := range input {
		copy(source[row+padding][padding:], input[row])
	}

	// make an output channel
	out := make(chan [][]rune)

	// generate kernels to the channel async
	go func() {
		defer close(out)

		// sweep the top left corner of the kernel, including the top/left pad, but excluding the bottom/right pad
		for y := range len(source) - padding {
			for x := range len(source[y]) - padding {

				// the kernel contents are size rows starting at y each comprising size elements starting at x
				kernel := make([][]rune, size)
				for row := range size {
					kernel[row] = source[y+row][x : x+size]
				}
				out <- kernel
			}
		}
	}()

	return out
}

// each element is a set of coordinates in the kernel used to build a check string
var chains = [][]int{
	// top row
	{0, 0, 0, 1, 0, 2, 0, 3},
	// left column
	{0, 0, 1, 0, 2, 0, 3, 0},
	// descending diagonal
	{0, 0, 1, 1, 2, 2, 3, 3},
	// ascending diagonal
	{0, 3, 1, 2, 2, 1, 3, 0},
}

func A() error {
	// read input into a grid
	grid := readGrid()

	// sweep a kernel over the grid
	kernels := kernel(grid, 4)

	// track the matches
	var matches int

	// process each kernel
	for kernel := range kernels {

		// for each chain to check in this kernel
		for _, chain := range chains {
			// for each pair of indices in the chain, extract the rune at that position and append it to build a check
			// string
			var check []rune
			for i := 0; i < len(chain)-1; i += 2 {
				check = append(check, kernel[chain[i]][chain[i+1]])
			}
			checkString := string(check)

			// if the check string matches the search (or its reverse) then count it
			if checkString == "XMAS" || checkString == "SAMX" {
				matches++
			}
		}
	}

	fmt.Println(matches)

	return nil
}

func B() error {
	// read input into a grid
	grid := readGrid()

	// sweep a kernel over the grid
	kernels := kernel(grid, 3)

	// track the matches
	var matches int

	// process each kernel
	for kernel := range kernels {

		// there's only one chain we need to check
		chain := []int{0, 0, 0, 2, 1, 1, 2, 0, 2, 2}

		// for each pair of indices in the chain, extract the rune at that position and append it to build a check
		// string
		var check []rune
		for i := 0; i < len(chain)-1; i += 2 {
			check = append(check, kernel[chain[i]][chain[i+1]])
		}
		checkString := string(check)

		// if the check string matches any valid pattern then count it
		switch checkString {
		case "MMASS", "MSAMS", "SSAMM", "SMASM":
			matches++
		}
	}

	fmt.Println(matches)

	return nil
}
