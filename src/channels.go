package main

import (
	"fmt"
	"sync"
)

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	// receiving goroutine
	go func() {
		// receiving data from the channel and assinging to the i variable
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	// sending goroutine
	go func() {
		// sending data from the channel
		ch <- 1337
		wg.Done()
	}()
	wg.Wait()

	// the arrow ' <- ' syntax send and receive value from the channel
}
