package _1_plutonian_pebbles

import (
	"aoc24/intmath"
	"aoc24/linkedlist"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func A() error {
	sentinel := &linkedlist.LinkedList[int]{-1, nil, nil}
	head := sentinel

	for _, field := range strings.Fields(strings.TrimSpace(input)) {
		stone, _ := strconv.Atoi(field)
		head = head.InsertAfter(stone)
	}

	for range 25 {
		for head = sentinel.Next; head != nil; head = head.Next {
			strVal := strconv.Itoa(head.Value)

			switch {
			case head.Value == 0:
				head.Value = 1
			case len(strVal)%2 == 0:
				half := len(strVal) / 2
				head.Value, _ = strconv.Atoi(strVal[:half])
				newVal, _ := strconv.Atoi(strVal[half:])
				head = head.InsertAfter(newVal)
			default:
				head.Value = head.Value * 2024
			}
		}
	}

	var length int
	for head = sentinel.Next; head != nil; head = head.Next {
		length++
	}

	fmt.Println(length)
	return nil
}

func B() error {
	memo := make(map[int]map[int]int)

	var stones int

	for _, field := range strings.Fields(strings.TrimSpace(input)) {
		stone, _ := strconv.Atoi(field)
		stones += stonesAt(stone, 75, memo)
	}

	fmt.Println(stones)
	return nil
}

// stonesAt calculates the number of stones a given stone will generate after some number of blinks.
// memo is a memoisation map of stone->(blinks->result) which is required for efficient recursion.
func stonesAt(stone, blinks int, memo map[int]map[int]int) (result int) {
	switch {
	case blinks == 0:
		return 1
	case memo[stone] != nil && memo[stone][blinks] != 0:
		return memo[stone][blinks]
	case stone == 0:
		result = stonesAt(1, blinks-1, memo)
	case intmath.Digits(stone)%2 == 0:
		strStone := strconv.Itoa(stone)
		half := len(strStone) / 2
		left, _ := strconv.Atoi(strStone[:half])
		right, _ := strconv.Atoi(strStone[half:])
		result = stonesAt(left, blinks-1, memo) + stonesAt(right, blinks-1, memo)
	default:
		result = stonesAt(stone*2024, blinks-1, memo)
	}
	stoneMemo, ok := memo[stone]
	if !ok {
		stoneMemo = make(map[int]int)
		memo[stone] = stoneMemo
	}
	stoneMemo[blinks] = result
	return
}
