package test

import (
	"bou.ke/monkey"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestProcessFirstLine(t *testing.T) {
	output := ProcessFirstLine()
	expectOutput := "line00"
	assert.Equal(t, expectOutput, output)
}

func TestFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	output := ProcessFirstLine()
	expectOutput := "line000"
	assert.Equal(t, expectOutput, output)
}
