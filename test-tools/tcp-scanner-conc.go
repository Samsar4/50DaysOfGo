package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// act as a synchronized counter
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		// increment the counter via wg.Add() each time goroutine is created
		wg.Add(1)
		go func(j int) {
			// decrements counter when one work has been performed
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	// main() calls wg.Wat() to blocks until all the work
	// has been done and counter has returned to zero
	wg.Wait()
}
