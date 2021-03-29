package main

import (
	"fmt"
	"sync"
)

// Syncronizing goroutines:

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished.
// At the same time, Wait can be used to block until all goroutines have finished.
// https://golang.org/pkg/sync/#WaitGroup

var wg = sync.WaitGroup{}

func main() {
	var msg = "Test message"
	// increment the WaitGroup
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		// done decrements the WaitGroup counter by one
		wg.Done()
	}(msg)

	// reassingning the value to check
	msg = "Other message"
	// wait blocks until the WaitGroup counter is zero
	wg.Wait()

}
