package main

import (
	"fmt"
)

// Testing some loops
// for is Go’s only looping construct

func main() {
	for a := 0; a < 5; a++ {
		fmt.Println(a)
	}

	fmt.Println("--------------------")

	for x, y := 0, 0; x < 5; x, y = x+1, y+2 {
		fmt.Println(x, y)
	}

	fmt.Println("--------------------")

	i := []int{}
	for j := 10; j > 0; j-- {
		i = append(i, j)
	}
	fmt.Println(i)
	fmt.Println(i[3])

	fmt.Println("--------------------")

}
