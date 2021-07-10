package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// Print command and the following args
	fmt.Printf("Command: %v Args: %v", os.Args[1], os.Args[2:])
	for i, arg := range os.Args {
		fmt.Printf("Arg: %s Index: %v\n", arg, i)
	}
}
