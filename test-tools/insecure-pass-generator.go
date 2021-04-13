package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// random output every time compiled
	rand.Seed(time.Now().UTC().UnixNano())

	// check for input
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./pass-generator <length of pass> \n %s \n", os.Args[0])
		os.Exit(1)
	}

	passLength := os.Args[1]
	// converts string arguments into int using strconv.Atoi
	if s, err := strconv.Atoi(passLength); err == nil {
		fmt.Println(randomStr(s))
	}

}

func randomStr(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(3400, 3450))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
