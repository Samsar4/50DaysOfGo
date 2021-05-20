package main

import (
	"fmt"
)

func FindInt(arr []int, f func(int) bool) (int, error) {
	for _, v := range arr {
		if f(v) {
			return v, nil
		}
	}
	return 0, fmt.Errorf("No matches found")
}
func FindFloat64(arr []float64, f func(float64) bool) (float64, error) {
	for _, v := range arr {
		if f(v) {
			return v, nil
		}
	}
	return 0.0, fmt.Errorf("No matches found")
}
func FindBool(arr []bool, f func(bool) bool) (bool, error) {
	for _, v := range arr {
		if f(v) {
			return v, nil
		}
	}
	return false, fmt.Errorf("No matches found")
}
func FindString(arr []string, f func(string) bool) (string, error) {
	for _, v := range arr {
		if f(v) {
			return v, nil
		}
	}
	return "", fmt.Errorf("No matches found")
}

func main() {
	fmt.Println(FindInt([]int{1, 1, 2}, func(x int) bool { return x%2 == 0 }))
	fmt.Println(FindFloat64([]float64{1.3, 3.4, 2.0}, func(x float64) bool { return false }))
}
