package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	bruteString := BruteForceJoin(os.Args)
	fmt.Printf("bruteString: %s", bruteString)

	fastString := FasterJoin(os.Args)
	fmt.Printf("fastString: %s", fastString)
}

// BruteForceJoin is a naive implementation of a string builder
func BruteForceJoin(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	// Print command and the following args
	fmt.Printf("Command: %v Args: %v", args[1], args[2:])
	for i, arg := range args {
		fmt.Printf("Arg: %s Index: %v\n", arg, i)
	}
	return s
}

// FasterJoin is a less naive implementation of a string builder
func FasterJoin(args []string) string {
	return strings.Join(args[1:], " ")
}
