package main

// formula (F to C): C = (F - 32) * 5/9

import "fmt"

func FahrenheitToCelsius(d float64) float64 {
	return (d - 32.0) * 5.0 / 9.0
}

func main() {

	fah := 100.00
	fmt.Printf("Fahrenheit: %g\nCelsius: %.2f\n", fah, FahrenheitToCelsius(fah))

}
