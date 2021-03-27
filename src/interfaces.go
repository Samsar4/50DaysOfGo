package main

import "fmt"

// Interfaces are named collections of method signatures.
// The interface defines the behavior for similar type of objects.
// Is declared using the type keyword, followed by the name of the interface and the keyword interface.

// note:
// 	-> interface defines BEHAVIOR for type of obj
//  -> structs defines DATA

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello yo"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
