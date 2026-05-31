package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWord(t *testing.T) {
	for _, tc := range testCases {
		result := occurencesCnc(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}

type testCase struct {
	description string
	input       []string
	expected    []int
}

var testCases = []testCase{
	{
		description: "no urls",
		input:       []string{},
		expected:    []int{},
	},
	{
		description: "one url",
		input:       []string{"https://go.dev/blog/intro-generics"},
		expected:    []int{27},
	},
	{
		description: "two urls",
		input:       []string{"https://go.dev/blog/intro-generics", "https://go.dev/blog/gofix"},
		expected:    []int{27, 3},
	},
}

func BenchmarkSeq(b *testing.B) {
	for b.Loop() {
		occurencesSeq(testCases[2].input)
	}
}

func BenchmarkCnc(b *testing.B) {
	for b.Loop() {
		occurencesCnc(testCases[2].input)
	}
}
