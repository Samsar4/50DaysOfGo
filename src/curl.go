package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// init is called before main. 13
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./curl <url>")
		os.Exit(-1)
	}
}

// main is the entry point for the application
func main() {
	// get a response from web server
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// copies from the body to Stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
