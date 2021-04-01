// Studying arrays and slices

// SLICES vs ARRAYS:
// Slices --> are more flexible, powerful and convenient to use in Go Lang
// Arrays --> fixed-length sequences of items of the same type; arrays in Go are values not pointers;
// Botttom line: most of use cases people use slices instead of arrays.

// More info: https://blog.golang.org/slices-intro

package main

// Multiple import with parenthesis
import (
	"fmt"
)

func main() {

	// You can declare array, the type, size and value in one line
	// set values
	set := [5]int{5, 6, 4, 4, 3}
	fmt.Printf("Set: %v \n", set)

	// or you can declare the array with N (as array grow) using '...'
	setN := [...]int{4, 1, 10, 4}
	fmt.Printf("set array length to N: %v \n", setN)
	fmt.Printf("----------------------------------\n")

	// adding strings on empty arrays
	var students [3]string
	students[0] = "Elliot"
	students[1] = "Sam"
	students[2] = "Jack"
	fmt.Printf("Your name is: %v \n", students)
	fmt.Printf("The size of the array of names is: %v \n", len(students))
	fmt.Printf("----------------------------------\n")

	// two dimensions array (matrix)
	var twoD [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Printf("Two-dimensional array: %v \n", twoD)
	fmt.Printf("----------------------------------\n")

	// copying arrays	in go is actually copying to a new array
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 10

	fmt.Printf("Original array: %v\n", a)
	fmt.Printf("Copied and Changed array: %v\n", b)

	// you can copy the array and pass the memory pointer using & (ampersand)
	// this results in both arrays pointing to same memory location (same value)
	x := [...]int{1, 2, 3}
	y := &x
	y[1] = 10
	x[0] = 20
	fmt.Printf("Two arrays pointing same memory location: %v \n %v\n", x, y)
	fmt.Printf("----------------------------------\n")

	// SLICES
	// slices is reference-based, they always point to the same memory pointer
	slice := []int{1, 2, 3}
	slice_b := slice
	slice[0] = 10
	fmt.Println(slice)
	fmt.Println(slice_b)

	fmt.Printf("----------------------------------\n")

	// slices operations
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl_1 := sl[:]   // slice of ALL elements
	sl_2 := sl[3:]  // slice from 4th element to end
	sl_3 := sl[3:6] // slice the 4th, 5th, and 6th elements
	sl_4 := sl[:6]  // slice first 6 elements

	// this change will affect every slice
	sl[1] = 10

	fmt.Println("slice ------------->", sl)
	fmt.Println("all elements [:] -->", sl_1)
	fmt.Println("[3:] -------------->", sl_2)
	fmt.Println("[3:6] ------------->", sl_3)
	fmt.Println("[:6] -------------->", sl_4)

	fmt.Printf("----------------------------------\n")

	// Using make() function to create slices
	sl_5 := make([]int, 3)
	fmt.Println(sl_5)
	fmt.Printf("Length: %v \n", len(sl_5))
	fmt.Printf("Capacity: %v \n", cap(sl_5))

	fmt.Printf("----------------------------------\n")

	// cap vs len
	// --> cap() tells you the capacity of the underlying array
	// --> len() tells you how many items are in the array (SIZE)
	// --> append() function to add elements to slice (may cause expensive copy operation if underlying array is too small)

	sl_6 := []int{}
	fmt.Println(sl_6)
	fmt.Printf("Length: %v \n", len(sl_6))
	fmt.Printf("Capacity: %v \n", cap(sl_6))
	fmt.Println("Added 3 elements to slice:")
	sl_6 = append(sl_6, 10, 20, 30)
	fmt.Println(sl_6)
	fmt.Printf("Length: %v \n", len(sl_6))
	fmt.Printf("Capacity: %v \n", cap(sl_6))

	fmt.Printf("----------------------------------\n")

	// testing the capacity limit
	sl_7 := make([]int, 0, 3) // slice w/ len = 0 and cap = 3
	fmt.Println("Array:", sl_7)

	for g := 0; g < 5; g++ {
		sl_7 = append(sl_7, g)
		fmt.Printf("Cap: %v | Len: %v | MemoryPointer: %p\n", cap(sl_7), len(sl_7), sl_7)
	}

	fmt.Println("Once the capacity is met, append() will return a new slice with a larger capacity!")

	fmt.Printf("----------------------------------\n")

}
