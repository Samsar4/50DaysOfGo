package main

import (
	"io"
	"log"
	"net"
)

// handler function that echo the received data
func echo(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)

	for {
		// Receive data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		// Sends data via conn.Write
		log.Println("Writing data...")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write the data")
		}
	}
}

func main() {
	// bind to TCP port 10080 on all interfaces
	listener, err := net.Listen("tcp", ":10080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:10080")
	for {
		// wait for connection; net.Conn is created when connection established
		// the infinite loop ensure that the server will continue to listen for connections even after one has been received
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// concurrent is called so that other connections don't block while waiting for the handler function to complete
		go echo(conn)
	}
}
