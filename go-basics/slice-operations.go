package main

import (
	"fmt"
)

func main() {

	// Adding values to the first position on slice
	slice := []int{0, 1, 3}
	sliceTwo := []int{4}
	// "..." ellipsis expand slice content
	slice = append(sliceTwo, slice...)
	fmt.Println(slice)

	// 2nd example
	s := []int{10, 20, 30}
	s = append([]int{40}, s...)
	fmt.Println(s)

	fmt.Println("---------------")

	// Removing values by slicing it
	sl := []int{10, 20, 30, 40, 50}
	sl = sl[1:]
	fmt.Println(sl)

	// Removing specific values
	ss := []int{100, 200, 300, 400, 500, 600}
	index := 2 // 300
	// appending {100,200} + {400,500...}
	ss = append(ss[:index], ss[index+1:]...)
	fmt.Println(ss)

	fmt.Println("---------------")

	// Copying slices using copy() func
	numbers := []int{1, 1, 2, 3, 5, 8}
	numbersCopied := make([]int, len(numbers))

	copy(numbersCopied, numbers)
	for i := range numbersCopied {
		numbersCopied[i] *= 5
	}

	fmt.Printf("numbers: %v %p \nnumbersCopied: %v %p \n", numbers, numbers, numbersCopied, numbersCopied)

	fmt.Println("---------------")

}
