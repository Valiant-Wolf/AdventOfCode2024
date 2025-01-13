package _8_resonant_collinearity

import (
	. "aoc24/util"
	. "aoc24/vect"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func A() error {
	width, height, antennae := parseMap(input)

	bounds := NewRect(Vector{}, Vector{width, height})

	antinodes := make(map[Vector]bool)

	for _, nodes := range antennae {
		for i, first := range nodes {
			for _, second := range nodes[i+1:] {
				difference := first.Sub(second)

				firstAntinode := first.Add(difference)
				if bounds.Contains(firstAntinode) {
					antinodes[firstAntinode] = true
				}

				secondAntinode := second.Sub(difference)
				if bounds.Contains(secondAntinode) {
					antinodes[secondAntinode] = true
				}
			}
		}
	}

	fmt.Println(len(antinodes))

	return nil
}

func B() error {
	width, height, antennae := parseMap(input)

	bounds := NewRect(Vector{}, Vector{width, height})

	antinodes := make(map[Vector]bool)

	for _, nodes := range antennae {
		for i, first := range nodes {
			for _, second := range nodes[i+1:] {
				difference := first.Sub(second)

				for checkNode := first; bounds.Contains(checkNode); checkNode = checkNode.Add(difference) {
					antinodes[checkNode] = true
				}

				for checkNode := second; bounds.Contains(checkNode); checkNode = checkNode.Sub(difference) {
					antinodes[checkNode] = true
				}
			}
		}
	}

	fmt.Println(len(antinodes))

	return nil
}

func parseMap(stringMap string) (width, height int, antennae map[rune][]Vector) {
	antennae = make(map[rune][]Vector)
	rows := strings.Split(strings.TrimSpace(stringMap), "\n")
	for y, row := range rows {
		for x, char := range []rune(row) {
			if char != '.' {
				antennae[char] = append(antennae[char], Vector{x, y})
			}
		}
	}
	width, height = len(rows[0]), len(rows)
	return
}
