package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	string1 := "Some secrets why not!"

	// generates the SHA-1 hash
	hashed := sha1.New()
	hashed.Write([]byte(string1))

	// checksum of the data
	bs := sha1.Sum(nil)

	fmt.Printf("#1 String: %s\n", string1)
	fmt.Printf("SHA-1 hash: %x\n", bs)

	string2 := "more secrets right?"
	h := md5.New()
	io.WriteString(h, string2)
	b := md5.Sum(nil)
	fmt.Printf("#2 String: %s\n", string2)
	fmt.Printf("MD5 hash: %x\n", b)

	// hashing a file example
	f, err := os.Open("assets/secret.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hh := md5.New()
	if _, err := io.Copy(hh, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File: %s\n MD5 hash: %x\n", f.Name(), hh.Sum(nil))
}
