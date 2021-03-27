package main

import (
	"fmt"
)

// Methods is a special type of function that executes in context of a type

func main() {
	s := sign{
		line:  "Welcome",
		line2: "Pal",
	}
	s.welcome()

}

type sign struct {
	line  string
	line2 string
}

func (s sign) welcome() {
	fmt.Println(s.line, s.line2)
}
