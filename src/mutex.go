package main

import (
	"fmt"
	"runtime"
	"sync"
)

// this example complete kills the concurrency and parallelism
// the mutexes forces to run on a single threaded way

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		// locking the mutexes on a single context
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Message #%v\n", counter)
	// unlock when goroutine is done
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
