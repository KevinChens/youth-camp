package test

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestJudgePassLineTre(t *testing.T) {
	output := JudgePassLine(70)
	expectOutput := true
	assert.Equal(t, expectOutput, output)
}

func TestJudgePassLineFalse(t *testing.T) {
	output := JudgePassLine(50)
	expectOutput := false
	assert.Equal(t, expectOutput, output)
}
