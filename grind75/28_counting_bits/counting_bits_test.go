package counting_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input    int
	expected []int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input:    2,
			expected: []int{0, 1, 1},
		},
		{
			input:    5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, countBits(test.input), "input: %d", test.input)
	}
}
