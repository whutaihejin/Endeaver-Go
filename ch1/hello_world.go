package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	x, y := 0, 0
	z := x + y
	fmt.Println("z = ", z)
	n := x + y
	fmt.Println("n = ", n)
}

// right declare
func X() {

}

// bad declare
/*
func X()
{

}*/
