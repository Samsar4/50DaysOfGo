package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

// struct type for JSON
type Msg struct {
	Greeting string
	Name     string
}

// struct type for XML; field tag inside colons, defines how data is handled
type Banner struct {
	Fizz string `xml:"id,attr"`      // attribute named tag
	Buzz string `xml:"parent>child"` // subelement of parent named child
}

func main() {
	m := Msg{"Hello", "Jason Bourne"}
	// Marshal() encodes struct to JSON, returning byte slice
	// n, _ := json.Marshal(m)
	// ---> Prints the sdout as JSON-encoded string representation of Msg struct
	// fmt.Println(string(n))

	prettyJSON, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		log.Fatal("Failed to generate JSON", err)
	}
	fmt.Printf("----------[JSON]-----------\n%s\n", string(prettyJSON))

	// Unmarshal() take the same byte slice and decoded it
	//json.Unmarshal(n, &m)

	// xml follows the same logic as json
	// XML encoder reflectively determines the names of elements
	// using tag directives, so each field is handled properly
	b := Banner{"coolIDbro", "buzz buzz"}
	//bx, _ := xml.Marshal(b)
	//fmt.Println(string(bx))

	prettyXML, err := xml.MarshalIndent(b, " ", "   ")
	if err != nil {
		log.Fatal("Failed to generate XML", err)
	}
	fmt.Printf("----------[XML]-----------\n%s\n", string(prettyXML))
}
