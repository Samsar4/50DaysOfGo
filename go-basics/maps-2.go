package main

import (
	"fmt"
)

type Car struct {
	model string
	year  int
}

func main() {
	cars := make(map[string]Car)

	cars["Audi"] = Car{"S4 Quattro", 2014}
	cars["Honda"] = Car{"Accord EX-L", 2018}
	cars["Peugeot"] = Car{"308 GT", 2017}
	cars["Tesla"] = Car{"Model S", 2014}

	for model, year := range cars {
		fmt.Println(model, year)
	}

	// lookup
	_, find := cars["Honda"]
	if find {
		fmt.Printf("Car found: %v\n", cars["Honda"])
	} else {
		fmt.Println("Car not found!")
	}

	// bool lookup #2
	_, ok := (cars["Tesla"])
	fmt.Printf("The key Tesla is there? --> %v\n", ok)
	_, ok = (cars["BMW"])
	fmt.Printf("The key BMW is there? --> %v\n", ok)
}
