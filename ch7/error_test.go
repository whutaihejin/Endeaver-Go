package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	assert.False(t, errors.New("EOF") == errors.New("EOF"))
}

type Errno uintptr // operating system error code

var errs = [...]string{
	1: "operation not permitted",   // EPERM
	2: "no such file or directory", // ENOENT
	3: "so such process",
}

func (e Errno) Error() string {
	if 0 <= int(e) && int(e) < len(errs) {
		return errs[e]
	}
	return fmt.Sprintf("errno %d", e)
}

func TestErrs(t *testing.T) {
	var err error = Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}
