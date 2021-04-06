package main

import (
	"fmt"
	"strings"
)

// Fibonacci series
func main() {

	f0 := 0
	f1 := -1
	f2 := 1
	var fibo []int

	for f0 < 100 {
		f0 = f1 + f2
		f1 = f2
		f2 = f0
		fibo = append(fibo, f0)
	}

	// Drawing Fibonacci with asterisks
	fmt.Println("Fibonacci asterisk draw: \n")
	i := 0
	for range fibo {
		x := []string{strings.Repeat("*", fibo[i])}
		fmt.Println(x)
		i++
	}

	fmt.Printf("Fibonacci slice: %v \n", fibo)

}
