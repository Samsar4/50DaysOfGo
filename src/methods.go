package main

import (
	"fmt"
)

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
