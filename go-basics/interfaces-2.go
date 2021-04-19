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

func (list *GenericList) RemoveFirst() interface{} {
	return list.RemoveIndex(0)
}

func (list *GenericList) RemoveLast() interface{} {
	return list.RemoveIndex(len(*list) - 1)
}

func main() {

	list := GenericList{1337, "Apple", 10, true, 20, "Car", 3.14, false}

	fmt.Printf("\nFULL List:\n%v\n\n", list)

	fmt.Printf("Removing the first element (%v)...\n%v\n", list.RemoveFirst(), list)
	fmt.Printf("Removing the last element (%v)...\n%v\n", list.RemoveLast(), list)
	fmt.Printf("Removing index 3 (%v)...\n%v\n", list.RemoveIndex(3), list)
	fmt.Printf("Removing index 0 (%v)...\n%v\n", list.RemoveIndex(0), list)
	fmt.Printf("Removing last index (%v)...\n%v\n", list.RemoveIndex(len(list)-1), list)

}
