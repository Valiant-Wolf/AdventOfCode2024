package _2_garden_groups

import (
	"aoc24/vect"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func A() error {
	plots, _, _ := findPlots(input)

	var cost int

	for _, plotList := range plots {
		for _, plot := range plotList {

			var area, perimeter int
			for pos := range plot {
				area++
				for _, direction := range cardinals {
					checkPos := pos.Add(direction)
					if _, same := plot[checkPos]; !same {
						perimeter++
					}
				}
			}

			cost += area * perimeter
		}
	}

	fmt.Println(cost)

	return nil
}

func B() error {
	plots, width, height := findPlots(input)

	var cost int

	// calculate cost for each plot
	for _, plotList := range plots {
		for _, plot := range plotList {

			area := len(plot)
			var sides int
			// for each vertical slice of the plot
			for _, row := range rows(plot, width, height) {
				tops := make([]int, 0)
				bottoms := make([]int, 0)
				// check each plant
				for plant := range row {
					// if there is no plant above, list it as a top
					if pos := plant.Add(vect.Up()); !plot[pos] {
						tops = append(tops, plant.X)
					}
					// no plant below -> bottom
					if pos := plant.Add(vect.Down()); !plot[pos] {
						bottoms = append(bottoms, plant.X)
					}
				}

				// sort the lists of plant positions
				slices.Sort(tops)
				slices.Sort(bottoms)
				// each run of consecutive plants is a distinct side
				sides += consecutive(tops) + consecutive(bottoms)
			}

			// repeat for horizontal slices
			for _, col := range cols(plot, width, height) {
				lefts := make([]int, 0)
				rights := make([]int, 0)
				for plant := range col {
					if pos := plant.Add(vect.Left()); !plot[pos] {
						lefts = append(lefts, plant.Y)
					}
					if pos := plant.Add(vect.Right()); !plot[pos] {
						rights = append(rights, plant.Y)
					}
				}

				slices.Sort(lefts)
				slices.Sort(rights)
				sides += consecutive(lefts) + consecutive(rights)
			}

			cost += area * sides
		}
	}

	fmt.Println(cost)
	return nil
}

var cardinals = []vect.Vector{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func findPlots(input string) (map[rune][]map[vect.Vector]bool, int, int) {
	inputRows := strings.Split(strings.TrimSpace(input), "\n")
	plantMap := make(map[rune]map[vect.Vector]bool)

	for y, row := range inputRows {
		for x, char := range []rune(row) {
			plant, ok := plantMap[char]
			if !ok {
				plant = make(map[vect.Vector]bool)
				plantMap[char] = plant
			}

			plant[vect.Vector{x, y}] = true
		}
	}

	height, width := len(inputRows), len(inputRows[0])

	plots := make(map[rune][]map[vect.Vector]bool)
	for char, openPlants := range plantMap {
		// openPlants is the set of all plants not yet assigned to a plot

		// foundPlots is the list of plots found for this plant
		foundPlots := make([]map[vect.Vector]bool, 0)

		for len(openPlants) != 0 {
			// create a new plot
			plot := make(map[vect.Vector]bool)
			foundPlots = append(foundPlots, plot)

			// activePlants contains a front of active plants that will move across a plot
			activePlants := make(map[vect.Vector]bool)
			// activePlants begins with one plant not in a plot
			for k := range openPlants {
				plot[k] = true
				activePlants[k] = true
				delete(openPlants, k)
				break
			}

			// while the front is still alive
			for len(activePlants) != 0 {
				// newPlants is a set to contain newly activated plants
				newPlants := make(map[vect.Vector]bool)
				// for each plant in  front
				for k := range activePlants {
					// check each cardinal direction
					for _, direction := range cardinals {
						checkPos := k.Add(direction)
						// check if there is an open plant there
						if _, open := openPlants[checkPos]; open {
							// if there is, add it to the plot, and the new front, and mark it as not open
							plot[checkPos] = true
							newPlants[checkPos] = true
							delete(openPlants, checkPos)
						}
					}
				}

				// make the new front active
				activePlants = newPlants
			}

			// at this point the front has moved through the whole plot;
			// if there are still open plants, this process is repeated until they are all assigned a plot
		}

		plots[char] = foundPlots
	}

	return plots, width, height
}

func rows(plots map[vect.Vector]bool, width, height int) []map[vect.Vector]bool {
	result := make([]map[vect.Vector]bool, 0)
	for y := range height {
		row := make(map[vect.Vector]bool)
		for x := range width {
			if pos := (vect.Vector{x, y}); plots[pos] {
				row[pos] = true
			}
		}
		if len(row) != 0 {
			result = append(result, row)
		}
	}
	return result
}

func cols(plots map[vect.Vector]bool, width, height int) []map[vect.Vector]bool {
	result := make([]map[vect.Vector]bool, 0)
	for x := range width {
		col := make(map[vect.Vector]bool)
		for y := range height {
			if pos := (vect.Vector{x, y}); plots[pos] {
				col[pos] = true
			}
		}
		if len(col) != 0 {
			result = append(result, col)
		}
	}
	return result
}

func consecutive(sequence []int) int {
	last := -2
	count := 0
	for _, entry := range sequence {
		if entry-last > 1 {
			count++
		}
		last = entry
	}
	return count
}
