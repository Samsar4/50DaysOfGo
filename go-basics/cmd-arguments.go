package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Error \n %s \n", os.Args)
		os.Exit(1)
	}

	var a string
	a = os.Args[1]

	if a == "command" {
		fmt.Println("Correct!")
	}

	os.Exit(0)
}
