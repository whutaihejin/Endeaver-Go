// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var dest = flag.String("d", "dest.dat", "-d dest file")

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	file, err := os.Open(*dest)
	if err != nil {
		fmt.Printf("%.2fs elapsed, fetchall open file: %v\n", time.Since(start).Seconds(), err)
		return
	}
	fmt.Fprintf(file, "%.2fs elapsed")
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	reply, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	n, err := io.Copy(ioutil.Discard, reply.Body)
	reply.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", seconds, n, url)
}
