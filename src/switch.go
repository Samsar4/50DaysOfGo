package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1:
		fmt.Println("Case 1 ok")
	case 2:
		fmt.Println("Case 2 ok")
	default:
		fmt.Println("Default case (every num instead 1 or 2 case)")
	}

	// you can add multiple tasks on a single case
	switch 3 {
	case 1, 5, 10:
		fmt.Println("one, five ten")
	case 2, 4, 6:
		fmt.Println("two for six")
	default:
		fmt.Println("another number")
	}

	// you can add comparsion and logical operators
	n := 10
	switch {
	case n%2 == 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}
}
