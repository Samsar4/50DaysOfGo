package main

import (
	"net/http"
)

// ---- Control Flow
// Defer, Panic and Recover:
// https://blog.golang.org/defer-panic-and-recover

// ---- Panic

// - Occur when program cannot continue at all
// - Don't use when file can't be opened, unless its critical
// - use for unrecovereable events (cannot obatain TCP port for webserver)
// - Function will stop executing (deferred functions will still fire)
// - If nothing handles panic, program will exit

func main() {
	// fmt.Println("1) First")
	// panic("something bad happened!!")
	// fmt.Println("3) Third")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Yoo, Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
