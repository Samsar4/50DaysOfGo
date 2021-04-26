package main

import (
	"fmt"
	"time"
)

func FibonacciGenerator(n int) func() {

	return func() {
		a, b := 0, 1

		fibo := func() int {
			a, b = b, a+b

			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fibo())
		}
	}

}

// calculate execution time
func Time(function func()) {
	start := time.Now()

	function()

	fmt.Printf(
		"\nExecution time: %s \n-------------------------\n",
		time.Since(start))
}

func main() {
	Time(FibonacciGenerator(6))
	Time(FibonacciGenerator(13))
	Time(FibonacciGenerator(28))
}
