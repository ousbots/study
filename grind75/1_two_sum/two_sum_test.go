package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums     []int
	target   int
	expected []int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {
		ans := twoSum(test.nums, test.target)
		assert.Equal(t, test.expected, ans)
	}
}

func TestAdditional(t *testing.T) {
	tests := []Test{
		{
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			nums:     []int{2, 9, 9, 9, 9, 9, 2},
			target:   4,
			expected: []int{0, 6},
		},
		{
			nums:     []int{-1, 9, 9, 9, 9, 9, 4},
			target:   3,
			expected: []int{0, 6},
		},
	}

	for _, test := range tests {
		ans := twoSum(test.nums, test.target)
		assert.Equal(t, test.expected, ans)
	}
}
