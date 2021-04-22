package main

import (
	"fmt"
	"html"
)

func init() {
	str := html.UnescapeString("&#128513;")
	fmt.Println(str)
}

func second() {
	str := html.UnescapeString("&#129300;")
	fmt.Println(str)
}

func third(msg string) {
	fmt.Printf("%s \n", msg)
}

func Conversion(euroConverted float64) (euro float64, converted float64, realBR float64) {
	euro = 6.78
	fee := 1.09
	feeConversion := 2.60
	wiseFee := fee + feeConversion

	realBR = euroConverted

	euroConverted -= wiseFee
	euroConverted /= euro

	return euro, euroConverted, realBR
}

func main() {
	// conversion call
	euro, converted, realBR := Conversion(100.00)

	fmt.Println("-------------------")
	fmt.Printf("Real: R$%.2f\n"+"Euro: €%.2f\n"+
		"Transfer: €%.2f\n",
		realBR, euro, converted)
	fmt.Println("-------------------")

	// emojis
	second()
	str := html.UnescapeString("&#128542;")
	third(str)
}
