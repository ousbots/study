package single_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums     []int
	expected int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, singleNumber(test.nums))
	}
}
