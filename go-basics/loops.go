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

	z := 0
	for ; z < 8; z++ {
		fmt.Println(z)
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
	fmt.Println(i[5])

	fmt.Println("--------------------")

	for o := 0; o <= 2; o++ {
		for p := 0; p <= 2; p++ {
			fmt.Println(o + p)
		}
	}

	fmt.Println("--------------------")

	col := []int{1, 2, 3}

	for k, v := range col {
		fmt.Println(k, v)
	}

	fmt.Println("--------------------")

	string1 := "loremp ipsum"
	for _, v := range string1 {
		//fmt.Println(k, v)         // unicode value
		fmt.Println(string(v)) // string value
	}

}
