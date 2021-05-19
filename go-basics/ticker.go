package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			// await the values as they arrive every 500ms
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick: ", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// once won't receive any more values it stops after 1600ms
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
