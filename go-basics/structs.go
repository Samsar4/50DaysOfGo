package main

import (
	"fmt"
	"reflect"
)

// Structs are typed collections of fields. They’re useful for grouping data together to form records; Collections of disparate data types that describe single concept

// Capitalized : Exported from the package (Age, Name, Books, Pets)
// Lower case : Internal to the package (age, name, books, pets)

type Person struct {
	age   int
	name  string
	books []string
	pets  []string
}

// No inheritance, but is possible to use composition via embedding:
type Animal struct {
	Name   string
	Origin string
}

type Cat struct {
	Animal   // <-- embedded struct
	Speed    float32
	EyeColor string
}

// struct with arbitrary value; using reflect
type Address struct {
	Street string `required max:"100"`
	Code   int
}

func main() {
	firstPerson := Person{
		age:   20,
		name:  "John Doe",
		books: []string{},
		pets: []string{
			"Black Cat",
			"White cat",
			"Gray cat",
		},
	}
	fmt.Println(firstPerson)

	fmt.Println("----------------------------")

	// Anonymous structs
	personTwo := struct{ name string }{name: "Jean Paul"}
	personThree := struct{}{}

	fmt.Println(personTwo, personThree)

	fmt.Println("----------------------------")

	// Animal and Cat
	b := Cat{}
	b.Name = "Balthazar"
	b.Origin = "Australia"
	b.Speed = 48.5
	b.EyeColor = "Yellow"

	fmt.Println(b.Name)
	fmt.Println(b)

	fmt.Println("----------------------------")

	theAnimal := Animal{Name: "Yomama", Origin: "South Africa"}
	fmt.Println(theAnimal)

	fmt.Println("----------------------------")

	theCat := Cat{}
	theCat.Name = "Felix"
	theCat.Origin = "France"
	theCat.Speed = 48.5
	theCat.EyeColor = "Blue"
	fmt.Println(theCat.Name)
	fmt.Println(theCat)

	fmt.Println("----------------------------")

	// Using reflect package
	// TypeOf extract dynamic type info;
	// FieldByName returns the struct field with given name
	t := reflect.TypeOf(Address{})
	field, _ := t.FieldByName("Street")
	fmt.Println(field.Tag)

}
