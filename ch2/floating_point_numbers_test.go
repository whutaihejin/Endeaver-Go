package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFloatingPointNumber(t *testing.T) {
	var f float32 = 16777216
	fmt.Println(f == f+1)
	assert.True(t, f == f+1)
	var n float64 = 16777216
	assert.False(t, n == n+1)
	const e = 2.71828 // approximately
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

	fmt.Println(compute(1))
	fmt.Println(compute(-1))
}

func compute(v int) (value float64, ok bool) {
	if v > 0 {
		return 0, false
	}
	return 1, true
}
