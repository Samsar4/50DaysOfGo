package main

import (
	"fmt"
	"net"
	"os"
)

// ResolveIPAddr -> returns an address of IP end point.
// If the host in the address parameter is not a literal IP address, ResolveIPAddr resolves the address to an address of IP end point. Otherwise, it parses the address as a literal IP address. The address parameter can use a host name, but this is not recommended, because it will return at most one of the host name's IP addresses.

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./resolve-ip-addr <url>\n %s \n ", os.Args[0])
		fmt.Println("Usage:", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Error!", err.Error())
		os.Exit(1)
	}

	fmt.Println("Resolved address: ", addr.String())
	os.Exit(0)
}
