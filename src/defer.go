package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ---- Control Flow
// Defer, Panic and Recover:
// https://blog.golang.org/defer-panic-and-recover

// ---- Defer
func main() {

	//	fmt.Println("1")
	//	defer fmt.Println("2")
	//	fmt.Println("3")
	// 	output: 1 3 2

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// closes the resource after error checking
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
