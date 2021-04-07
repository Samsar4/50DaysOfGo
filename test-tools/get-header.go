package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// Establishes a TCP connection, send the request string, read and print the response

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./get-header <url>:<port> \n %s \n ")
		os.Exit(1)
	}

	service := os.Args[1]

	// ResolveTCPAddr resolves the address to an address of TCP end point
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	//  DialTCP acts like Dial for TCP networks.
	// If laddr is nil, a local address is automatically chosen.
	// If the IP field of raddr is nil or an unspecified IP address, the local system is assumed.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error! \n %s", err.Error())
		os.Exit(1)
	}
}
