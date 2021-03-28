package main

import "fmt"

// Interfaces are named collections of method signatures.
// The interface defines the behavior for similar type of objects.
// Is declared using the type keyword, followed by the name of the interface and the keyword interface.

// Note:
// 	-> interface defines BEHAVIOR for type of obj
//  -> structs defines DATA

// Best practices:
// * Use many, small interfaces (single method interfaces are some of the most powerful and flexible [io.Writer, io.Reader, inteface{} ])
// * DON'T export interfaces for types that will be consumed
// * DO export interfaces for types that will be used by package

func main() {
	// 1st example
	var w Writer = ConsoleWriter{}
	w.Write([]byte("1st example: Hello yo"))

	// 2nd example
	Inteiro := IntCounter(0)
	var inc Incrementer = &Inteiro
	fmt.Println("2nd example (counter):")
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

// 1st example
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// 2nd example
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
