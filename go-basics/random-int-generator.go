package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// The default number generator is deterministic, so it’ll produce the same sequence of numbers each time by default. To produce varying sequences, give it a seed that changes.
	// Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.
	rand.Seed(time.Now().UnixNano())
}

func main() {
	n := 0

	for {
		n++
		// rand.Intn -> a non-negative pseudo-random number in [0,n) from the default Source.
		i := rand.Intn(100000)
		fmt.Println(i)
		// infinite loop until breaks
		if i%3 == 0 {
			break
		}
	}
	fmt.Printf("-----------\nIterations: %d \n", n)
}
