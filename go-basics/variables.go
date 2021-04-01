// This is my first go program to test basic & abstract concepts 

package main

import "fmt"

// Block for variables
var (
	a string = "Primeira string"
	b string = "segunda string"
	c int = 500
)

func main() {
	fmt.Println("\nPrinting variables:")
	fmt.Println("-------------------")
	var d int = 42
	x := "Yo"
	fmt.Println(d,x,a,b,c)
	yo()
	stringz()
	constants()
}

func yo() {
	// comparing 42 to 1 direct on assignment
	i := 42 == 1
	fmt.Printf("%v %T\n", i, i)
}

func stringz() {
	fmt.Println("\nPrinting strings and runes:")
	fmt.Println("-------------------")
	// strings are immutable , double quote " - UTF-8
	s := "This is a string. "

	// You can convert to a collections of bytes (slice of byte)
	// to make easy for send data to other app for example
	b := []byte(s)
	fmt.Println(b)

	// runes (single quote ') - enconded UTF-32 , alias for int32
	var r rune = 'a'
	fmt.Printf("%v %T\n", r, r)
}

func constants() {
	// constants are immutable but can be shadowed
	fmt.Println("\nPrinting constants:")
	fmt.Println("-------------------")
	const myConst int = 50
	const myConst2 string = "pineapple"
	const myConst3 float32 = 3.14
	const myConst4 bool = true
	// enumerate constants:
	// -> iota keyword representes successive int const (0,1,2...) PER BLOCK
	// note: the value of iota resets for each block 
	// note2: enum expressions -> arithmetic, bitwise operations, bitshifting

	const (
		bb = iota + 40
		bb2 = iota
		bb3 = iota
	)

	fmt.Printf("%v %T\n", myConst, myConst)
	fmt.Printf("%v %T\n", myConst2, myConst2)
	fmt.Printf("%v %T\n", myConst3, myConst3)
	fmt.Printf("%v %T\n", myConst4, myConst4)
	fmt.Printf("%v \n", bb)
	fmt.Printf("%v \n", bb2)
	fmt.Printf("%v \n", bb3)
}
