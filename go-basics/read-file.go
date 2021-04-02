package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Read file to a string

func main() {
	content, err := ioutil.ReadFile("assets/lorem.txt")
	if err != nil {
		log.Fatal(err)
	}
	// convert []byte to string
	text := string(content)
	fmt.Printf("The file contains:\n\n %s", text)
}
