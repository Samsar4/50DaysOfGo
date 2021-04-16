package main

import (
	"fmt"
	"sort"
)

func main() {

	squares := make(map[int]int, 10)

	for n := 1; n <= 10; n++ {
		squares[n] = n * n
	}
	fmt.Println(squares)

	// slice func make([]type, cap, len) []type
	numbers := make([]int, 0, len(squares))

	for n := range squares {
		numbers = append(numbers, n)
	}

	// sort.Ints() sorts a slice of ints in increasing order
	sort.Ints(numbers)

	for _, num := range numbers {
		squared := squares[num]
		fmt.Printf("%d^2 -> %d\n", num, squared)
	}

}
