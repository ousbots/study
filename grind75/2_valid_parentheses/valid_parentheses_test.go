package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input    string
	expected bool
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input:    "()",
			expected: true,
		},
		{
			input:    "()[]{}",
			expected: true,
		},
		{
			input:    "(]",
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isValid(test.input))
	}
}

func TestAdditional(t *testing.T) {
	tests := []Test{
		{
			input:    ")",
			expected: false,
		},
		{
			input:    "(",
			expected: false,
		},
		{
			input:    "((()))",
			expected: true,
		},
		{
			input:    "([{[()]}])",
			expected: true,
		},
		{
			input:    "([{)]}",
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isValid(test.input))
	}
}
