package _7_bridge_repair

import (
	"aoc24/util"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func A() error {
	var result int

	// for each line in the input
	for _, equation := range strings.Split(strings.TrimSpace(input), "\n") {
		// extract the equation value and numbers
		numbers, _ := util.Sliceatoi(strings.Fields(strings.ReplaceAll(equation, ":", "")))
		value, numbers := numbers[0], numbers[1:]

		// recursively check if the equation is valid
		valid := validate(value, numbers, numbers[0], 1)

		if valid {
			result += value
		}
	}

	fmt.Println(result)

	return nil
}

func B() error {
	var result int

	// for each line in the input
	for _, equation := range strings.Split(strings.TrimSpace(input), "\n") {
		// extract the equation value and numbers
		numbers, _ := util.Sliceatoi(strings.Fields(strings.ReplaceAll(equation, ":", "")))
		value, numbers := numbers[0], numbers[1:]

		// recursively check if the equation is valid
		valid := concatValidate(value, numbers, numbers[0], 1)

		if valid {
			result += value
		}
	}

	fmt.Println(result)

	return nil
}

func validate(value int, numbers []int, current int, depth int) bool {
	// short-circuit if overshot
	if current > value {
		return false
	}

	// if we're at depth, check if we hit the value
	if depth == len(numbers) {
		return current == value
	}

	// calculate options
	add := current + numbers[depth]
	mul := current * numbers[depth]

	// recurse both branches
	return validate(value, numbers, add, depth+1) || validate(value, numbers, mul, depth+1)
}

func concatValidate(value int, numbers []int, current int, depth int) bool {
	// short-circuit if overshot
	if current > value {
		return false
	}

	// if we're at depth, check if we hit the value
	if depth == len(numbers) {
		return current == value
	}

	// calculate options
	add := current + numbers[depth]
	mul := current * numbers[depth]
	con, _ := strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(numbers[depth]))

	// recurse all branches
	return concatValidate(value, numbers, add, depth+1) || concatValidate(value, numbers, mul, depth+1) || concatValidate(value, numbers, con, depth+1)
}
