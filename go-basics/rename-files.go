package main

import (
	"log"
	"os"
)

func main() {
	src := "/home/ubuntu/100DaysOfGo/go-basics/assets/renamed-files/3.txt"
	dst := "/home/ubuntu/100DaysOfGo/go-basics/assets/renamed-files/3-renamed.txt"

	err := os.Rename(src, dst)
	if err != nil {
		log.Fatal(err)
	}

}
