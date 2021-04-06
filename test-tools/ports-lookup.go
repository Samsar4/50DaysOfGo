package main

import (
	"fmt"
	"net"
	"os"
)

// Return the port number by given service name
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Example: ./port-lookup tcp ssh \n %s \n", os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	service := os.Args[2]
	// LookupPort looks up the port for the given network and service.
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error! ", err.Error())
		os.Exit(2)
	}
	fmt.Println("Service port ", port)
	os.Exit(0)
}
