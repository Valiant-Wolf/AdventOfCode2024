package main

import (
	"aoc24/challenges"
	"fmt"
	"os"
)

func main() {

	cmdStr := os.Args[1]

	cmd := challenges.Challenges[cmdStr]
	if cmd == nil {
		fmt.Printf("'%s' is not a valid command\n", cmdStr)
		return
	}

	err := cmd()
	if err != nil {
		fmt.Println(err)
	}
}
