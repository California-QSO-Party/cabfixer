package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin_AllPositiveNumbers(t *testing.T) {
	r := min(4, 7)
	assert.Equal(t, 4, r, `Expected 4 as minumum of 4 and 7`)
}

func TestMin_OneNegativeNumber(t *testing.T) {
	r := min(-14, 7)
	assert.Equal(t, -14, r, `Expected -14 as minumum of -14 and 7`)
}
