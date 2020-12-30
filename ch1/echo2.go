// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	a := ""
	var b string
	var c = ""
	var d string = ""
	fmt.Println(a, b, c, d)
}
