package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestSatisfaction(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	// cannot use time.Second (type time.Duration) as type io.Writer in assignment:
	//	time.Duration does not implement io.Writer (missing Write method)
	// w = time.Second
	fmt.Printf("%T\n", w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// cannot use new(bytes.Buffer) (type *bytes.Buffer) as type io.ReadWriteCloser in assignment:
	//	*bytes.Buffer does not implement io.ReadWriteCloser (missing Close method)
	// rwc = new(bytes.Buffer)
	fmt.Printf("%T\n", rwc)

	w = rwc
	// cannot use w (type io.Writer) as type io.ReadWriteCloser in assignment:
	//	io.Writer does not implement io.ReadWriteCloser (missing Close method)
	// rwc = w
}

type IntSet struct {

}

func (*IntSet) String() string {
	return "x"
}

func TestIntSet(t *testing.T) {
	// cannot call pointer method on IntSet literal
	//./interfaces_interface_satisfaction_test.go:42:18: cannot take the address of IntSet literal
	// var _ = IntSet{}.String()
	var s IntSet
	var _ = s.String()
	var _ fmt.Stringer = &s
	// cannot use s (type IntSet) as type fmt.Stringer in assignment:
	//	IntSet does not implement fmt.Stringer (String method has pointer receiver)
	// var _ fmt.Stringer = s
}

func TestSatisfy(t *testing.T) {
	os.Stdout.Write([]byte("hello"))
	os.Stdout.Close()

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello"))
	//  w.Close undefined (type io.Writer has no field or method Close)
	// w.Close()
}