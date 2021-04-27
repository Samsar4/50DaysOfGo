package main

import (
	"fmt"
)

func chanelling(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	// closing the channel to avoid deadlock error
	close(c)
}

func main() {

	// buffer size 3
	c := make(chan int, 3)

	go chanelling(c)

	// check if the value was assigned or not, avoiding zero value in this case
	for {
		value, ok := <-c
		if ok {
			fmt.Println(value)
		} else {
			break
		}
	}

}
