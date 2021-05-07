package main

import (
	"fmt"
	"log"
	"os"
)

// defines an io.Reader to read from Stdin
type FooReader struct{}

// READS the data from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in>")
	return os.Stdin.Read(b)
}

// defines an io.Writer to write to Stdout
type FooWriter struct{}

// WRITES the data to Stdout
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	// instantiate reader & writer
	var (
		reader FooReader
		writer FooWriter
	)
	// buffer to hold input and output
	input := make([]byte, 4096)

	// Using reader to read the input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Cannot read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Using writer to write to output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Cannot write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
