package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error) {
	return 0, nil
}

func Printf(format string, args ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, args...)
}

func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.String()
}

// Writer is the interface that wraps the basic Write method.
type MyWriter interface {
	// Write writes len(p) bytes from p to the underlying data stream.
	// It returns the number of bytes written from p(0 <= n <= len(p))
	// and any error encountered that caused the write to stop early.
	// Write must return a non-nil error if it returns n < len(p).
	// Write must not modify the slice data, even temporarily.

	// Implementations must no retain p.
	MyWrite(p []byte) (n int, err error)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("hello"))
	assert.Equal(t, ByteCounter(5), c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	assert.Equal(t, ByteCounter(12), c)
}
