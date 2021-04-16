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
	countryCodes := map[string]int{
		"Afghanistan":    93,
		"Albania":        355,
		"Algeria":        213,
		"American Samoa": 684,
		"Andorra":        376,
		"Angola":         244,
	}
	fmt.Println(countryCodes)

	// using make() function
	new_map := make(map[string]float64)
	new_map["Apple"] = 3.14
	fmt.Println(new_map)

	// Acessing some Values by the Key
	fmt.Printf("The country code of Algeria is: %v \n", countryCodes["Algeria"])

	// Adding some Key and Values to the map
	countryCodes["Australia"] = 61
	fmt.Printf("The country code of Australia is: %v \n", countryCodes["Australia"])

	// The map
	fmt.Println(countryCodes)

	// Deleting key & value by key
	delete(countryCodes, "Andorra")
	fmt.Println(countryCodes)

	// Check if the value exists using '_' - anom var
	_, ok := (countryCodes["Afghanistan"])
	fmt.Printf("The value exist? --> %v \n", ok)

	// When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next. If you require a stable iteration order you must maintain a separate data structure that specifies that order.

	var keys []string
	for k := range countryCodes {
		keys = append(keys, k)
	}

	fmt.Println(keys)

}
