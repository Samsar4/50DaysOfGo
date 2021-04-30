package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	input := "If you look for perfection, you'll never be content."
	encoded := b64.StdEncoding.EncodeToString([]byte(input))
	fmt.Printf("Base 64 Encoded: %s\n", encoded)

	decode, err := b64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decode error!", err)
		return
	}
	fmt.Printf("Base 64 Decoded: %s\n", decode)

	path := "SomeSecret0123456789!@#$"
	url := "https://example.org/api/"

	urlEncoded := b64.URLEncoding.EncodeToString([]byte(path))
	fmt.Printf("URL: %s%s\n", url, urlEncoded)

	urlDecode, err := b64.URLEncoding.DecodeString(urlEncoded)

	if err != nil {
		fmt.Println("URL Decode error!", err)
		return
	}
	fmt.Printf("URL Base 64 Decoded: %s%s\n", url, urlDecode)

}
