package longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	strs     []string
	expected string
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			strs:     []string{},
			expected: "",
		},
		{
			strs:     []string{"farfignewton"},
			expected: "farfignewton",
		},
		{
			strs:     []string{"abcd", "abc", "ab"},
			expected: "ab",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, longestCommonPrefix(test.strs))
	}
}
