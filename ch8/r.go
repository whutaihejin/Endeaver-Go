package main

import "fmt"

func main() {
	fmt.Println("\ra\rb\rc")
	fmt.Println("\ra\bb\bc")
	s := "hello world"
	for _, v := range s {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	for _, v := range s {
		fmt.Printf("\b%c", v)
	}
	fmt.Println()
}
