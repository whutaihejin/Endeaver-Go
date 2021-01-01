package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu2 sync.Mutex
var count2 int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler2(w http.ResponseWriter, r *http.Request) {
	mu2.Lock()
	count2++
	mu2.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter2(w http.ResponseWriter, r *http.Request) {
	var n int
	mu2.Lock()
	n = count2
	mu2.Unlock()
	fmt.Fprintf(w, "Count %d\n", n)
}
