package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	num      int
	expected bool
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			num:      121,
			expected: true,
		},
		{
			num:      -121,
			expected: false,
		},
		{
			num:      10,
			expected: false,
		},
		{
			num:      1,
			expected: true,
		},
		{
			num:      111,
			expected: true,
		},
		{
			num:      1219121,
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isPalindrome(test.num))
	}
}
