package main

import (
	"fmt"
)

// Branching with if and else in Go is straight-forward.
// You can have an if statement without an else
// You don't need parentheses around conditions in Go

func main() {
	num := 10
	numTwo := 100

	if num > 1 && numTwo > 1 {
		if num <= numTwo {
			fmt.Printf("%v is less than %v \n", num, numTwo)
		}
		if num > numTwo {
			fmt.Printf("%v is greater than %v \n", num, numTwo)
		}
		if numTwo > num {
			fmt.Printf("%v is greater than %v \n", numTwo, num)
		}
		if num < numTwo || num == numTwo {
			fmt.Printf("%v is less OR equal to %v \n", num, numTwo)
		}
		if num < numTwo && num != numTwo {
			fmt.Printf("%v is less AND not %v \n", num, numTwo)
		}
	}

	// if + else
	par := 8
	if par%2 == 0 {
		fmt.Printf("%v is even \n", par)
	} else {
		fmt.Printf("%v is odd \n", par)
	}

	// if + else if
	if n := 3.14; n == 3 {
		fmt.Println("ALMOST PI")
	} else if n == 3.14 {
		fmt.Println("PI")
	} else {
		fmt.Println("NO PI")
	}
}
