package main

import (
	"fmt"
	"time"
)

// Goroutines are functions or methods that run concurrently with other functions or methods.
// Goroutines can be thought of as light weight threads.
// The cost of creating a Goroutine is tiny when compared to a thread. Hence it's common for Go applications to have thousands of Goroutines running concurrently.

func message(str string) {
	for i := 0; i <= 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str, i)
	}
}

func main() {
	go message("Yo!")
	message("Yo2")
}
