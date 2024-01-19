package valid_anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	first    string
	second   string
	expected bool
}

func TestGiven(t *testing.T) {
	tests := []Test{
		{
			first:    "anagram",
			second:   "nagaram",
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isAnagram(test.first, test.second))
	}
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			first:    "",
			second:   "",
			expected: true,
		},
		{
			first:    "abccba",
			second:   "aabbcc",
			expected: true,
		},
		{
			first:    "abcccba",
			second:   "aabbcc",
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isAnagram(test.first, test.second))
	}
}
