package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan string)
	wg.Add(2)

	// Receive ONLY (data from the channel)
	go func(ch <-chan string) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// Send ONLY (data from the channel)
	go func(ch chan<- string) {
		ch <- "Message received :)"
		wg.Done()
	}(ch)
	wg.Wait()
}
