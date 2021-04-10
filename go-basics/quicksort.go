package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entryData := os.Args[1:]
	numbers := make([]int, len(entryData))

	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: ./quicksort <num1> <num2> <num3> ... \n %s \n", os.Args)
		os.Exit(1)
	}

	for i, n := range entryData {
		// strcov.Atoi() converts into integers
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("Error! %s is not a valid number \n", n)
			os.Exit(1)
		}
		numbers[i] = num

	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	lesser, greater := separate(n, pivot)

	return append(
		append(quicksort(lesser), pivot),
		quicksort(greater)...)

}

func separate(numbers []int, pivot int) (lesser []int, greater []int) {
	for _, n := range numbers {
		if n <= pivot {
			lesser = append(lesser, n)
		} else {
			greater = append(greater, n)
		}
	}
	return lesser, greater
}
