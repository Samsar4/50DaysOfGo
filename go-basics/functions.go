package main

import (
	"fmt"
	"regexp"
	"strings"
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

	// Anonymous functions
	func() {
		fmt.Println("#Anonymous! #1")
	}()

	anon := func() {
		fmt.Println("#Anonymous! #2")
	}
	anon()

	var fun func() = func() {
		fmt.Println("#Anonymous! #3")
	}
	fun()

	// using regexp, strings pkg and anonymous func to convert strings to uppercase
	// regex: "\b\w" --> grab the first letter on every string and aditional "\" escapes
	expr := regexp.MustCompile("\\b\\w")
	name := "alan mathison turing"

	// passing anonymous func inside the ReplaceAllStringFunc()
	process := expr.ReplaceAllStringFunc(
		name,
		func(s string) string {
			return strings.ToUpper(s)
		})

	fmt.Println(process)
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

// Can return addresses of local variables; Automatically promoted from local memory (stack) to shared memory (heap)
// "..." -> take the last arguments and return into a slice named values
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
