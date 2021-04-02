package main

import (
	"io/ioutil"
	"log"
)

// writes data to a file; if the file does not exist, WriteFile creates with permissions
func main() {
	quote := []byte("You can't trust code that you did not totally create yourself.")
	err := ioutil.WriteFile("assets/quote.txt", quote, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
