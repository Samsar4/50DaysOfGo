package main

import (
	"fmt"
)

func filter(nums []int, odd, even chan<- int, ready chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			even <- n
		} else {
			odd <- n
		}
	}
	ready <- true
}

func main() {
	odd, even := make(chan int), make(chan int)
	ready := make(chan bool)
	nums := []int{3, 337, 20, 16, 80, 13, 3459, 2, 50, 8, 9, 10, 99, 155}

	go filter(nums, odd, even, ready)

	var sliceOdd, sliceEven []int
	end := false

	for !end {
		select {
		case n := <-odd:
			sliceOdd = append(sliceOdd, n)
		case n := <-even:
			sliceEven = append(sliceEven, n)
		case end = <-ready:
		}
	}
	fmt.Printf("Odd: %v \nEven: %v\n", sliceOdd, sliceEven)
}
