package ch2

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestArithmetic(t *testing.T) {
	mask := uint32(0xffff)
	assert.Equal(t, uint32(0x08), mask&(1<<3))
	// the remainder operator %
	assert.Equal(t, -2, -5%3)
	assert.Equal(t, -2, -5%-3)
	assert.Equal(t, 1.25, 5.0/4.0)
	assert.Equal(t, 1, 5/4)
	assert.Equal(t, 1.25, 5*1.0/4)
	assert.Equal(t, 1.25, float64(5)/float64(4))

	// overflow
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	assert.Equal(t, uint8(0), u+1)
	assert.Equal(t, uint8(1), u*u)
	fmt.Printf("%x %#[1]x %#[1]X %[1]d %[1]b\n", 255*255)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	assert.Equal(t, int8(-128), i+1)
	assert.Equal(t, int8(1), i*i)
	fmt.Printf("%x, %[1]d, %[1]b\n", 127*127)

	//
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println("offset->", i)
		}
	}
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	var off int = 2
	assert.Equal(t, 4, 1<<off)

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	// infinite loop
	/*
		for i := uint32(2); i >= 0; i-- {
			fmt.Println(i)
		}*/

	/*
		var apples int32 = 1
		var oranges int16 = 1
		// invalid operation: apples + oranges (mismatched types int32 and int16)
		var compote int = apples + oranges // compile error
	*/
	{
		f := 3.141
		i := int(f)
		assert.Equal(t, 3.141, f)
		assert.Equal(t, 3, i)
		f = 1.99
		assert.Equal(t, 1, int(f))
		f = 1e100
		i = int(f)
		assert.Equal(t, 1e100, f)
		assert.Equal(t, -9223372036854775808, i) // result is implementation-dependent
		fmt.Println(f, i)
	}

	{
		o := 0666
		fmt.Printf("%d %[1]o, %#[1]o\n", o)
		x := int64(0xdeadbeef)
		fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	}

	{
		ascii := 'a'
		unicode := '国'
		newline := '\n'
		fmt.Printf("%d %[1]c %[1]q\n", ascii)
		fmt.Printf("%d %[1]c %[1]q\n", unicode)
		fmt.Printf("%d %[1]q\n", newline)
	}
}
