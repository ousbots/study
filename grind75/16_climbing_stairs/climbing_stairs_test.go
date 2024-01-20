package climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input    int
	expected int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input:    2,
			expected: 2,
		},
		{
			input:    3,
			expected: 3,
		},
		{
			input:    5,
			expected: 8,
		},
		{
			input:    45,
			expected: 1836311903,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, climbStairs(test.input))
	}
}
