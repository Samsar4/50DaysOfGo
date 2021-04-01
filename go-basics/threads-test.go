package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}
