package _5_print_queue

import (
	"aoc24/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func A() error {
	splitInput := strings.Split(strings.TrimSpace(input), "\n\n")

	rules := compileRules(splitInput[0])

	var correct int

	// for each line in the input
	for _, update := range strings.Split(splitInput[1], "\n") {

		// split the line into fields (separated by commas)
		pages := strings.Split(update, ",")

		// convert the fields to ints
		intPages, err := util.Sliceatoi(pages)
		if err != nil {
			return err
		}

		// validate the update and add the middle page
		if validate(intPages, rules) {
			correct += intPages[len(intPages)/2]
		}
	}

	fmt.Println(correct)

	return nil
}

func B() error {
	splitInput := strings.Split(strings.TrimSpace(input), "\n\n")

	rules := compileRules(splitInput[0])

	var incorrect int

	// for each line in the input
	for _, update := range strings.Split(splitInput[1], "\n") {

		// split the line into fields (separated by commas)
		pages := strings.Split(update, ",")

		// convert the fields to ints
		intPages, err := util.Sliceatoi(pages)
		if err != nil {
			return err
		}

		// if the update is invalid
		if !validateOrSwap(intPages, rules) {
			// swap until it is valid to find the correct ordering
			for !validateOrSwap(intPages, rules) {
			}

			// add the middle page
			incorrect += intPages[len(intPages)/2]
		}
	}

	fmt.Println(incorrect)

	return nil
}

func validate(pages []int, rules map[int][]int) bool {
	seen := make(map[int]bool)

	// check each page in order
	for _, page := range pages {
		// mark this page as seen
		seen[page] = true

		// find all pages this page must not occur after (continue if none)
		rule, ok := rules[page]
		if !ok {
			continue
		}

		// for each page that must not occur before this one, fail if we've seen it
		for _, check := range rule {
			if seen[check] {
				return false
			}
		}
	}

	return true
}

func validateOrSwap(pages []int, rules map[int][]int) bool {
	seen := make(map[int]bool)

	// check each page in order
	for i, page := range pages {
		// mark this page as seen
		seen[page] = true

		// find all pages this page must not occur after (continue if none)
		rule, ok := rules[page]
		if !ok {
			continue
		}

		// for each page that must not occur before this one
		for _, check := range rule {

			// if we've seen that page before
			if seen[check] {

				// find that page
				j := 0
				for pages[j] != check {
					j++
				}

				// swap this page with that one
				pages[i], pages[j] = pages[j], pages[i]

				return false
			}
		}
	}

	return true
}

// creates a map of rules that can be read as "<key> must not occur after <values>"
func compileRules(input string) map[int][]int {
	rules := make(map[int][]int)

	// for each line in the input
	for _, ruleString := range strings.Split(input, "\n") {
		// extract the fields
		fields, _ := util.Sliceatoi(strings.Split(ruleString, "|"))

		// append or create the new rule entry
		rule, _ := rules[fields[0]]
		rules[fields[0]] = append(rule, fields[1])
	}

	return rules
}
