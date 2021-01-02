package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	fmt.Println(string(1234567))
	r := []rune(string(1234567))
	assert.Equal(t, '\uFFFD', r[0])
	assert.Equal(t, int32('\uFFFD'), r[0])
	// assert.Equal(t, int('\uFFFD'), r[0]) // expected: int(65533) actual  : int32(65533)
	assert.Equal(t, "京", string(0x4eac))
	assert.Equal(t, "A", string(65))

	s := "程序"
	fmt.Printf("% x\n", s)
	r = []rune(s)
	fmt.Printf("% x\n", r)
	n := utf8.RuneCountInString(s)
	fmt.Println(n)

	{
		// first method
		n = 0
		for _, _ = range s {
			n++
		}
		assert.Equal(t, 2, n)
		// second method
		n = 0
		for range s {
			n++
		}
		assert.Equal(t, 2, n)
		// third method
		assert.Equal(t, 2, utf8.RuneCountInString(s))
		assert.Equal(t, 2, CountRune(s))
	}

	{
		str := "hello, 世界"
		for i, r := range str {
			fmt.Printf("%d\t%q\t%d\n", i, r, r)
		}
	}

	{
		s := "Hello, 世界"
		assert.Equal(t, 13, len(s))
		assert.Equal(t, 9, utf8.RuneCountInString(s))
	}

}

func TestStringToNumber(t *testing.T) {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	assert.Equal(t, "123", y)
	assert.Equal(t, "123", strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(strconv.FormatInt(int64(x), 8))
	fmt.Println(strconv.FormatInt(int64(x), 16))
	assert.Equal(t, "1111011", fmt.Sprintf("%b", x))
	assert.Equal(t, "1111011", strconv.FormatInt(int64(x), 2))

	// string to int
	x, err := strconv.Atoi("123")
	assert.Nil(t, err)
	assert.Equal(t, 123, x)
	x, err = strconv.Atoi("123nm")
	assert.NotNil(t, err)
	assert.Equal(t, 0, x)

	m, err := strconv.ParseInt("123", 10, 64)
	assert.Nil(t, err)
	assert.Equal(t, int64(123), m)
	m, err = strconv.ParseInt("123mn", 10, 64)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), m)
}

func CountRune(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
		c++
	}
	return c
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func TestInts(t *testing.T) {
	values := []int{1, 2, 3}
	assert.Equal(t, "[1, 2, 3]", intsToString(values))
}

func TestStrings(t *testing.T) {
	s := "hello, world hello"
	assert.True(t, strings.Contains(s, "hello"))
	assert.Equal(t, 2, strings.Count(s, "ll"))
	assert.Equal(t, 2, strings.Count(s, "hello"))
	fmt.Println(strings.Fields(s))
	// var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}
	// Fields splits the string s around each instance of one or more consecutive white space
	// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
	// empty slice if s contains only white space.

	// Fields splits the string s around each instance of one or more consecutive white space
	// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
	// empty slice if s contains only white space.
	fmt.Println(strings.Fields("a\tb"))
	assert.Equal(t, []string{"a", "b"}, strings.Fields("a\tb"))
	assert.Equal(t, []string{"hello,", "world", "hello"}, strings.Fields(s))

	assert.True(t, strings.HasPrefix(s, "hell"))
	assert.False(t, strings.HasPrefix(s, "xx"))
	assert.Equal(t, 1, strings.Index(s, "ell"))
	assert.Equal(t, -1, strings.Index(s, "xx"))
	strArr := []string{"hello", "my", "dog"}
	assert.Equal(t, "hello my dog", strings.Join(strArr, " "))
	assert.Equal(t, "hello,my,dog", strings.Join(strArr, ","))
}
