package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x * x + y * y)
}

func TestHypot(t *testing.T) {
	assert.Equal(t, 5.0, hypot(3, 4))
}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return z
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

func TestFuncType(t *testing.T) {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
	assert.Equal(t, "func(int, int) int", fmt.Sprintf("%T", add))
	assert.Equal(t, "func(int, int) int", fmt.Sprintf("%T", sub))
	assert.Equal(t, "func(int, int) int", fmt.Sprintf("%T", first))
	assert.Equal(t, "func(int, int) int", fmt.Sprintf("%T", zero))
}
