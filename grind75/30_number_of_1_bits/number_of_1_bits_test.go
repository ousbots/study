package number_of_1_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	num      uint32
	expected int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			num:      11,
			expected: 3,
		},
		{
			num:      128,
			expected: 1,
		},
		{
			num:      4294967293,
			expected: 31,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, hammingWeight(test.num))
	}
}
