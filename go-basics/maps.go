package main

import (
	"fmt"
)

// Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function

// Note: Manipulating maps is memory reference based, be aware!
// Note2: Map can receive arrays not slices

// More info: https://blog.golang.org/maps

func main() {
	// new_map := map[keyType]valueType
	zipcodes := map[string]int{
		"Alcabideche":          2645,
		"Cascais":              2750,
		"Carcavelos":           2775,
		"Estoril":              2765,
		"Parede":               2775,
		"São Domingos de Rana": 2785,
	}
	fmt.Println(zipcodes)

	// using make() function
	new_map := make(map[string]int)
	fmt.Println(new_map)

	// Acessing some Values by the Key
	fmt.Printf("The first 4 digits zipcode of Cascais is: %v \n", zipcodes["Cascais"])

	// Adding some Key and Values to the map
	zipcodes["Apelação"] = 2680
	fmt.Printf("The first 4 digits zipcode of Apelação is: %v \n", zipcodes["Apelação"])

	// The map
	fmt.Println(zipcodes)

	// Deleting some Keys
	delete(zipcodes, "São Domingos de Rana")

	// Check if the value exists using ('_' anom var)
	_, ok := (zipcodes["Cascais"])
	fmt.Printf("The value exist? --> %v \n", ok)

	// When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next. If you require a stable iteration order you must maintain a separate data structure that specifies that order.

	var keys []string
	for k := range zipcodes {
		keys = append(keys, k)
	}

	fmt.Println(keys)

}
