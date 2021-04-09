package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: <host>:<port>\n %s \n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("some data"))
	checkError(err)

	var buf [1024]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))
	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error! \n %s", err.Error())
		os.Exit(1)
	}
}
