package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// port 1201 nucleus-sand
	service := ":1201"
	// ResolveTCPAddr -> returns an address of TCP end point
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// ListenTCP -> acts like Listen for TCP networks
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// Goroutine client handler; accepts multi threaded connections
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Closes the connection on exit
	defer conn.Close()
	var buf [1024]byte

	for {
		// READ up to 512 bytes
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		// WRITE the n bytes read
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error!! \n %s \n", err.Error())
		os.Exit(1)
	}
}
