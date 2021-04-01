package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1:
		fmt.Println("Case 1 ok")
	case 2:
		fmt.Println("Case 2 ok")
	default:
		fmt.Println("Default case (every num instead 1 or 2 case)")
	}

	// you can add multiple tasks on a single case
	switch 3 {
	case 1, 5, 10:
		fmt.Println("one, five ten")
	case 2, 4, 6:
		fmt.Println("two for six")
	default:
		fmt.Println("another number")
	}

	// you can add comparsion and logical operators
	n := 10
	switch {
	case n%2 == 0:
		fmt.Printf("%v is Even\n", n)
	default:
		fmt.Printf("%v is Odd\n", n)
	}

	// type swtiches
	var i interface{} = 3.14
	switch i.(type) {
	case int:
		fmt.Printf("%v is int\n", i)
	case string:
		fmt.Printf("%v is string\n", i)
	case float64:
		fmt.Printf("%v is float64\n", i)
	default:
		fmt.Printf("%v is another value\n", i)
	}

}
