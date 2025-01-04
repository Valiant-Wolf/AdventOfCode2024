package _3_mull_it_over

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

var simpleExpr = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var complexExpr = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)

func A() error {
	// read the input into a slice of slices
	// each inner slice has format [fullMatch left right]
	matches := simpleExpr.FindAllStringSubmatch(input, -1)

	var total int

	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])

		total += left * right
	}

	fmt.Println(total)

	return nil
}

func B() error {
	// read the input into a slice of slices
	// each inner slice has format [fullMatch left? right?]
	matches := complexExpr.FindAllStringSubmatch(input, -1)

	var total int
	enable := true

	for _, match := range matches {
		switch {
		case match[0] == "do()":
			enable = true
		case match[0] == "don't()":
			enable = false
		default:
			if enable {
				left, _ := strconv.Atoi(match[1])
				right, _ := strconv.Atoi(match[2])

				total += left * right
			}
		}
	}

	fmt.Println(total)

	return nil
}
