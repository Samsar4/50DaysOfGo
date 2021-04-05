package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./ip-mask <dotted_ip_address>\n %s \n ", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address!!")
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println("Address: ", addr.String(),
		"| Default mask length: ", bits,
		"| Leading ones count is ", ones,
		"| Mask (hexadecimal): ", mask.String(),
		"| Network is: ", network.String())
	os.Exit(0)
}
