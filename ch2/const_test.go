package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConst(t *testing.T) {
	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)
	fmt.Println((f - 32) * 5 / 9.0)
	fmt.Printf("%T %[1]v\n", (f-32)*5/9.0)
	fmt.Printf("%T %[1]v\n", 5/9*(f-32))
	fmt.Printf("%T %[1]v\n", 5.0/9*(f-32))
	assert.Equal(t, 100.0, 5.0/9*(f-32))
	assert.Equal(t, 100.0, 5/9.0*(f-32))
	assert.Equal(t, float64(0), 5/9*(f-32))
}

func TestConstVar(t *testing.T) {
	fmt.Printf("%T\n", 0)
	assert.Equal(t, "int", fmt.Sprintf("%T", 0))
	assert.Equal(t, "float64", fmt.Sprintf("%T", 0.0))
	assert.Equal(t, "complex128", fmt.Sprintf("%T", 0i))
	assert.Equal(t, "int32", fmt.Sprintf("%T", '\000'))
	assert.Equal(t, "int32", fmt.Sprintf("%T", '\u0000'))
	var i = int8(0)
	var ii int8 = 0
	fmt.Printf("%v %v\n", i, ii)
	{
		i := 0
		r := '\000'
		f := 0.0
		c := 0i
		assert.Equal(t, "int", fmt.Sprintf("%T", i))
		assert.Equal(t, "float64", fmt.Sprintf("%T", f))
		assert.Equal(t, "complex128", fmt.Sprintf("%T", c))
		assert.Equal(t, "int32", fmt.Sprintf("%T", r))
	}

	const (
		deadbeef = 0xdeedbeef
		a        = uint32(deadbeef)
		b        = float32(deadbeef)
		c        = float64(deadbeef)
		// d        = int32(deadbeef) // constant 3740122863 overflows int32
		// e        = float64(1e309) // constant 1e+309 overflows float64
		// f        = uint(-1) // constant -1 overflows uint
	)
}
