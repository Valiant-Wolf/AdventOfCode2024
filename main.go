package main

import (
	"fmt"
	"os"
)

var Commands = make(map[string]func() error)

func main() {

	cmdStr := os.Args[1]

	cmd := Commands[cmdStr]
	if cmd == nil {
		fmt.Printf("'%s' is not a valid command\n", cmdStr)
		return
	}

	err := cmd()
	if err != nil {
		fmt.Println(err)
	}
}
