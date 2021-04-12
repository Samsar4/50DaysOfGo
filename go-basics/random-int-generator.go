package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed uses the provided seed value to initialize the generator to a deterministic state; unix timestamp (miliseconds) int64
	// this will provide random output every time is compiled
	rand.Seed(time.Now().UnixNano())
	n := 0

	for {
		n++
		// rand.Intn -> a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
		i := rand.Intn(100000)
		fmt.Println(i)
		// infinite loop until breaks
		if i%3 == 0 {
			break
		}
	}
	fmt.Printf("-----------\nIterations: %d \n", n)
}
