package main

import (
	"fmt"
	"io"
)

// Reader is the interface that wraps the basic Read method.
// type Reader interface {
//   Read(p []byte) (n int, err error)
// }

type StringReader struct{}

func (l StringReader) Read(p []byte) (int, error) {
	p[0] = 'F'
	p[1] = 'U'
	p[2] = 'Z'
	p[3] = 'Z'
	p[4] = 'Y'

	return len(p), nil
}

// receives io.Reader as argument, call the method Read()
// convert slice of bytes into string
// return the string
func converter(r io.Reader) string {
	p := make([]byte, 5)
	r.Read(p)
	return string(p)
}

func main() {
	reader := StringReader{}
	fmt.Println(converter(reader))
}
