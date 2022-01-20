package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//[1, 1, 1, 1]  # => "b4 b4 b4 b4"
//[1, 0, 1, 1]  # => "b2 b4 b4"
//[1, 0]        # => "b1"

//[1,0,1,0,1,0,1,0] # => "b4 b4 b4 b4"
//[1,0,0,0,1,0,0,0] # => "b2 b2"
//[1,0,1,1] # => "b2 b4 b4"
//[1,0,0,0] # => "b1"
//[1, 0, 0, 0, 1, 1, 1, 1] # => "b2 b8 b8 b8 b8"

func TestConvert(t *testing.T) {
	tests := []struct {
		input    []int
		expected []string
	}{
		{
			input:    []int{1, 1, 1, 1},
			expected: []string{"b4", "b4", "b4", "b4"},
		},
		{
			input:    []int{1, 0, 1, 1},
			expected: []string{"b2", "b4", "b4"},
		},
		{
			input:    []int{1, 0},
			expected: []string{"b1"},
		},
		{
			input:    []int{1, 0, 1, 0, 1, 0, 1, 0},
			expected: []string{"b4", "b4", "b4", "b4"},
		},
		{
			input:    []int{1, 0, 0, 0, 1, 0, 0, 0},
			expected: []string{"b2", "b2"},
		},
		{
			input:    []int{1, 0, 1, 1},
			expected: []string{"b2", "b4", "b4"},
		},
		{
			input:    []int{1, 0, 0, 0},
			expected: []string{"b1"},
		},
		{
			input:    []int{1, 0, 0, 0, 1, 1, 1, 1},
			expected: []string{"b2", "b8", "b8", "b8", "b8"},
		},
	}

	for _, test := range tests {
		test := test

		result := convert(test.input)
		assert.Equal(t, result, test.expected)
	}
}
