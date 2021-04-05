package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip address\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	// The type IP from net package is defined as byte slices ( type IP []byte )
	// ParseIP() parses name(string) as dotted IP address (IPv4 or IPv6)
	address := net.ParseIP(name)
	if address == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", address.String())
	}
	os.Exit(0)
}
