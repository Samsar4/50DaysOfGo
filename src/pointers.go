package main

import (
	"fmt"
)

// Go supports pointers, allowing you to pass references to values and records within your program.

func main() {
	// both variables is pointing to the same memory location
	var x int = 50
	var y *int = &x
	fmt.Println(&x, y)
	fmt.Println(x, *y)
	*y = 30
	fmt.Println(x, *y)

	fmt.Println("----------------------")

	fmt.Println("type something to repeat:")
	var sequence string
	fmt.Scan(&sequence)

	for i := 1; i <= 5; i++ {
		fmt.Printf("%v) %s \n", i, sequence)
	}

}
