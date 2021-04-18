package main

import (
	"fmt"
)

type GenericList []interface{}

func (list *GenericList) RemoveIndex(index int) interface{} {

	l := *list
	removed := l[index]
	*list = append(l[0:index], l[index+1:]...)
	return removed
}

func (list *GenericList) RemoveStart() interface{} {
	return list.RemoveIndex(0)
}

func (list *GenericList) RemoveEnd() interface{} {
	return list.RemoveIndex(len(*list) - 1)
}

func main() {
	list := GenericList{
		1337, "Apple", 10, true, 20, "Car", 3.14, false,
	}
	fmt.Printf("Original list:\n%v\n\n", list)

	fmt.Printf("Removing from start, v:(%v), after: \n%v\n", list.RemoveStart(), list)
	fmt.Printf("Removing from end v:(%v), after: \n%v\n", list.RemoveEnd(), list)
	fmt.Printf("Removing from index 3: (%v), after:\n%v\n", list.RemoveIndex(3), list)
	fmt.Printf("Removing from index 0: (%v), after:\n%v\n", list.RemoveIndex(0), list)
	fmt.Printf("Removing last index: (%v), after:\n%v\n", list.RemoveIndex(len(list)-1), list)

}
