package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Join concatenates the elements of its first argument to create a single string.

func Join(params ...interface{}) string {
	// get the array or slice
	arr := reflect.ValueOf(params[0])
	// get the separator string
	sp := reflect.ValueOf(params[1]).String()

	ars := make([]string, arr.Len())

	for i := 0; i < arr.Len(); i++ {
		ars[i] = fmt.Sprintf("%v", arr.Index(i))
	}

	return strings.Join(ars, sp)

}

func main() {
	fmt.Println(Join([]int{3, 1, 3, 3, 7}, "~"))
}
