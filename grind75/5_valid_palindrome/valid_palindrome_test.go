package valid_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	word     string
	expected bool
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			word:     "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			word:     "race a car",
			expected: false,
		},
		{
			word:     " ",
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isPalindrome(test.word), "expected %t for %v", test.expected, test.word)
	}
}

func TestMore(t *testing.T) {
	tests := []Test{
		{
			word:     "",
			expected: true,
		},
		{
			word:     "abccba",
			expected: true,
		},
		{
			word:     "abcba",
			expected: true,
		},
		{
			word:     "1/a:2&b*b(2)a-1",
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isPalindrome(test.word), "expected %t for %v", test.expected, test.word)
	}
}
