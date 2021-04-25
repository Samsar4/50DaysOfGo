package main

import (
	"fmt"
)

func main() {
	a, b := 0, 1

	fibo := func() int {
		a, b = b, a+b
		return a
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n", fibo())
	}
}
