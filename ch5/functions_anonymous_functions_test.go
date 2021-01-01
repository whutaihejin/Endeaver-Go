package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestAnonymousFunc(t *testing.T) {
	v := strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
	assert.Equal(t, "IBM.:111", v)
}

// what the fuck?
// squares returns a function that returns the next
// square number each time it is called.
func Squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func TestSquare(t *testing.T) {
	f := Squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println("===========")
	s := Squares()
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
}

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string {
	"algorithms" : {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	// visitAll := func(items []string) { // ./functions_anonymous_functions_test.go:67:5: undefined: visitAll
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func TestTopoSort(t *testing.T) {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i + 1, course)
	}
}


