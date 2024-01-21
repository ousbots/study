package move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums     []int
	expected []int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
		{
			nums:     []int{0, 1, 0},
			expected: []int{1, 0, 0},
		},
		{
			nums:     []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
		{
			nums:     []int{1, 0, 0},
			expected: []int{1, 0, 0},
		},
	}

	for _, test := range tests {
		moveZeroes(test.nums)
		assert.Equal(t, test.expected, test.nums, "nums: %v", test.nums)
	}
}
