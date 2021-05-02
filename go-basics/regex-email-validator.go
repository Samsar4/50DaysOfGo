package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func main() {
	// Fake domain
	if e := "user13@some-made-up-domainn222.com"; !isEmailValid(e) {
		fmt.Println(e + " --> is not a valid email!")
	}
	// Fake domain
	if e := "test@a4904cvb4feafax4z.co.uk"; !isEmailValid(e) {
		fmt.Println(e + " --> not a valid email")
	}

	// Real domain
	if e := "test@amazon.co.uk"; !isEmailValid(e) {
		fmt.Println(e + " --> not a valid email")
	} else {
		fmt.Println(e + " --> valid email address!")
	}
}

// isEmailValid function checks th email structure and length test
// net.LookupMX checks the domain has a valid MX record.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}
