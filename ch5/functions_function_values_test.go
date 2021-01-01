package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func TestFuncValues(t *testing.T) {
	f := square
	fmt.Println(f(3))
	assert.Equal(t, 9, f(3))
	f = negative
	assert.Equal(t, -3, f(3))
	fmt.Printf("%T\n", f)
	// f = product
	// cannot use product (type func(int, int) int) as type func(int) int in assignment
}

func TestPanic(t *testing.T) {
	var f func(int) int
	fmt.Printf("%T\n", f)
	if f != nil {
		f(3)
	} else {
		fmt.Println("f func value is nil")
	}
	// f(3) // panic: call of nil function
	// panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x12984f7]
}

func add1(r rune) rune { return r + 1 }

func TestStringMap(t *testing.T) {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	assert.Equal(t, "IBM.:111", strings.Map(add1, "HAL-9000"))
	//
	v2 := strings.Map(add1, "VMS")
	assert.Equal(t, "WNT", v2)
	//
	v3 := strings.Map(add1, "Admix")
	assert.Equal(t, "Benjy", v3)
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted a n. Both functions are optional. pre is called
// before the children are visited (pre order) and post is called after(post order).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth * 2 , " ", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth * 2, " ", n.Data)
	}
}

func TestForEachNode(t *testing.T) {
	url := "https://golang.org"
	fetch := func(url string) (io.Reader, error) {
		reply, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		if reply.StatusCode != http.StatusOK {
			reply.Body.Close()
			return nil, fmt.Errorf("getting %s: %s", url, reply.Status)
		}
		return reply.Body, nil
	}
	r, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch %s error %s\n", url, err.Error())
		os.Exit(1)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error %s\n", err.Error())
		os.Exit(2)
	}
	forEachNode(doc, startElement, endElement)
}