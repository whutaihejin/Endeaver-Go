package main

import (
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
