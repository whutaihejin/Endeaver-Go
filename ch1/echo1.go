// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		// s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// for statement
	k := 1
	for k < 10 {
		fmt.Print(k, " ")
		k++
	}
	fmt.Println()

	// a traditional infinite loop
	/*
		for {

		}*/
}
