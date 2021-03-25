package main

import (
	"fmt"
)

// Go supports pointers, allowing you to pass references to values and records within your program.
// --> Pointers types use an asterisk(*) as a prefix to type pointed to [ex: *int]
// --> Use the addressof operator (&) to get address of variable

// * Note: All assignment operations in Go are copy operations;
// * Slices and maps contain internal pointers, so copies point to same underlying data

func zerop(val *int) {
	*val = 0
}

func zero(val2 int) {
	val2 = 0
}

func main() {
	fmt.Println("----------------------")

	num := 1
	fmt.Println("Initial value of num:", num)

	zero(num)
	fmt.Println("Zero func value:", num)

	zerop(&num)
	fmt.Println("Zerop (with int pointer) value:", num)
	fmt.Println("Zerop (&: mem addr):", &num)

	fmt.Println("----------------------")

	// both variables is pointing to the same memory location
	var x int = 50
	var y *int = &x
	fmt.Println(&x, y)
	fmt.Println(x, *y)
	*y = 30
	fmt.Println(x, *y)

	fmt.Println("----------------------")

	fmt.Println("type something to repeat:")
	var seq string
	// '&' gives the memory address of seq string
	fmt.Scan(&seq)

	for i := 1; i <= 5; i++ {
		fmt.Printf("%v - %s \n", i, seq)
	}

	fmt.Println("----------------------")

	arr := [3]int{1, 2, 3}
	arr0 := &arr[0] // pointer to an integer
	arr1 := &arr[1] // pointer to an integer
	fmt.Printf("%v %p %p \n", arr, arr0, arr1)

}
