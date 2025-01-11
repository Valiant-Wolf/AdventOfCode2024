package _0_hoof_it

import (
	. "aoc24/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func A() error {

	inputRows := strings.Split(strings.TrimSpace(input), "\n")
	heightMap := make([][]int, len(inputRows))

	for y, row := range inputRows {
		heightMap[y] = make([]int, len(row))
		for x, char := range []rune(row) {
			heightMap[y][x] = int(char - '0')
		}
	}

	bounds := NewRect(Vector{0, 0}, Vector{len(heightMap[0]), len(heightMap)})
	topo := &TopoMap{heightMap, bounds}

	var score int

	for y, row := range heightMap {
		for x, cell := range row {
			if cell == 0 {
				trailHead := Vector{x, y}
				summits := make(map[Vector]bool)
				distinctHike(0, trailHead, topo, summits)
				score += len(summits)
			}
		}
	}

	fmt.Println(score)

	return nil
}

func B() error {

	inputRows := strings.Split(strings.TrimSpace(input), "\n")
	heightMap := make([][]int, len(inputRows))

	for y, row := range inputRows {
		heightMap[y] = make([]int, len(row))
		for x, char := range []rune(row) {
			heightMap[y][x] = int(char - '0')
		}
	}

	bounds := NewRect(Vector{0, 0}, Vector{len(heightMap[0]), len(heightMap)})
	topo := &TopoMap{heightMap, bounds}

	var score int

	for y, row := range heightMap {
		for x, cell := range row {
			if cell == 0 {
				trailHead := Vector{x, y}
				score += hike(0, trailHead, topo)
			}
		}
	}

	fmt.Println(score)

	return nil
}

var cardinals = []Vector{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type TopoMap struct {
	heightMap [][]int
	bounds    Rect
}

func (t TopoMap) HeightAt(v Vector) (height int) {
	height = -1
	if t.bounds.Contains(v) {
		height = t.heightMap[v.Y][v.X]
	}
	return
}

func hike(height int, position Vector, topo *TopoMap) (score int) {
	if height == 9 {
		return 1
	}

	for _, direction := range cardinals {
		testPos := position.Add(direction)
		if topo.HeightAt(testPos) == height+1 {
			score += hike(height+1, testPos, topo)
		}
	}
	return
}

func distinctHike(height int, position Vector, topo *TopoMap, summits map[Vector]bool) {
	if height == 9 {
		summits[position] = true
		return
	}

	for _, direction := range cardinals {
		testPos := position.Add(direction)
		if topo.HeightAt(testPos) == height+1 {
			distinctHike(height+1, testPos, topo, summits)
		}
	}
}
