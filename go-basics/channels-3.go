package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}
	wg.Done()
}

// main is the entry point
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	for i := 1; i <= 10; i++ {
		// Send 10 integers on the channel.
		c <- i
	}

	close(c)
	wg.Wait()
}
