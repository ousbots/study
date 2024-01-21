package missing_number

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
			nums:     []int{3, 0, 1},
			expected: 2,
		},
		{
			nums:     []int{0, 1},
			expected: 2,
		},
		{
			nums:     []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, missingNumber(test.nums), "nums: %v", test.nums)
	}
}
