package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: args <value> <num>")
		os.Exit(1)
	}

	// prints the last arg on the slice os.Args
	foo := os.Args[len(os.Args)-1]
	fmt.Println(foo)
}
