package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// issue ping command on 1.1.1.1
	cmdOutput, err := exec.Command("ping", "-c 3", "1.1.1.1").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", cmdOutput)
}
