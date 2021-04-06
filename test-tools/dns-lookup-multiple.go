package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./dns-lookup-multiple <url> \n %s \n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	// LookupHost looks up the given host using the local resolver.
	// returns a slice of host's addresses.
	addresses, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	for _, s := range addresses {
		fmt.Println(s)
	}
	os.Exit(0)
}
