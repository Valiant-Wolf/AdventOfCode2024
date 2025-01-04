package _2_red_nosed_reports

import (
	"aoc24/intmath"
	"aoc24/util"
	_ "embed"
	"fmt"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

func A() (err error) {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var safe int

	defer func() {
		if caught, ok := recover().(error); ok {
			err = caught
		}
	}()

	// for each line in the input
	for _, report := range strings.Split(strings.TrimSpace(input), "\n") {

		// spin off a new goroutine
		wg.Add(1)
		go func(report string) {
			defer wg.Done()

			// split the line into fields (separated by whitespace)
			levels := strings.Fields(report)

			// convert the fields to ints
			intLevels, subErr := util.Sliceatoi(levels)
			if subErr != nil {
				panic(subErr)
			}

			// validate the report and increment if safe
			if validate(intLevels) {
				mu.Lock()
				defer mu.Unlock()

				safe++
			}
		}(report)
	}

	// wait for all goroutines to finish
	wg.Wait()

	fmt.Println(safe)

	return nil
}

func B() (err error) {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var safe int

	defer func() {
		if caught, ok := recover().(error); ok {
			err = caught
		}
	}()

	// define a function to concurrently increment the safe count
	safeFunc := func() {
		mu.Lock()
		defer mu.Unlock()
		safe++
	}

	// for each line in the input
	for _, report := range strings.Split(strings.TrimSpace(input), "\n") {

		// spin off a new goroutine
		wg.Add(1)
		go func(report string) {
			defer wg.Done()

			// split the line into fields (separated by whitespace)
			levels := strings.Fields(report)

			// convert the fields to ints
			intLevels, subErr := util.Sliceatoi(levels)
			if subErr != nil {
				panic(subErr)
			}

			// validate the whole report
			if validate(intLevels) {
				safeFunc()
				return
			}

			// validate each combination of the report with one element dropped
			for i := range intLevels {
				dropped := dropOne(intLevels, i)
				if validate(dropped) {
					safeFunc()
					return
				}
			}
		}(report)
	}

	// wait for all goroutines to finish
	wg.Wait()

	fmt.Println(safe)

	return nil
}

func validate(report []int) bool {
	var signs int

	for i, v := range report[:len(report)-1] {
		w := report[i+1]

		diff := w - v
		absDiff := intmath.Abs(diff)

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		signs += intmath.Sign(diff)
	}

	return intmath.Abs(signs) == len(report)-1
}

// dropOne creates a copy of a slice dropping the element at the specified index
func dropOne(slice []int, index int) []int {
	left, right := slice[:index], slice[index+1:]
	result := make([]int, len(left), len(slice)-1)
	_ = copy(result, left)
	result = append(result, right...)
	return result
}
