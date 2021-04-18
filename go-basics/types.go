package main

import (
	"fmt"
)

// type conversion example

type ShoppingList []string

func printSlice(slice []string) {
	fmt.Println("Slice -->", slice)
}

func printShoppingList(list ShoppingList) {
	fmt.Println("Shopping list -->", list)
}

func main() {

	list := ShoppingList{"Apple", "Chicken", "Olives"}
	slice := []string{"Apple", "Chicken", "Olives"}

	printSlice([]string(list))
	printShoppingList(ShoppingList(slice))

}
