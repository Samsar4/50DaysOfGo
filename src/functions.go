package main

import (
	"fmt"
)

// -> Functions in Go:
// * Basic Syntax, * Parameters, * Return values, * Anonymous functions, * Function as types and * Methods

func main() {

	greeting := "Hello Main"
	sayGreeting(&greeting)
	fmt.Println(greeting)

	fmt.Println("---------------------")

	for j := 1; j <= 2; j++ {
		sayMessage(j, "Message received!")
	}

	fmt.Println("---------------------")

	s := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("The sum is:", s)

	fmt.Println("---------------------")

	// anonymous functions
	func() {
		fmt.Println("#Anonymous!")
	}()

	f := func() {
		fmt.Println("#Anonymous! #2")
	}
	f()

	var fun func() = func() {
		fmt.Println("#Anonymous! #3")
	}
	fun()

}

func sayMessage(index int, msg string) {
	fmt.Println(index)
	fmt.Println(msg)
}

func sayGreeting(greeting *string) {
	fmt.Println(*greeting)
	*greeting = "Hello other function"
	fmt.Println(*greeting)
}

// "..." -> take the last arguments and return into a slice named values
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
