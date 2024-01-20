package longest_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input    string
	expected int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "abccccdd",
			expected: 7,
		},
		{
			input:    "a",
			expected: 1,
		},
		{
			input:    "abcdefghijjkkllmmmmmmmmmm",
			expected: 17,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, longestPalindrome(test.input))
	}
}
