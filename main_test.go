package main

import (
	"bytes"
	"os"
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

func TestProcessFile_FullTable(t *testing.T) {
	ProcessFile("a.log")
	assert.True(t, EqualFiles("a_answer.xcbr", "a.xcbr"))
}

func TestProcessFile_MissingFields(t *testing.T) {
	ProcessFile("b.log")
	assert.True(t, EqualFiles("b_answer.xcbr", "b.xcbr"))
}

func EqualFiles(f1, f2 string) bool {
	f1Content, err := os.ReadFile(f1)
	if err != nil {
		return false
	}
	f2Content, err := os.ReadFile(f2)
	if err != nil {
		return false
	}
	return bytes.Equal(f1Content, f2Content)
}
