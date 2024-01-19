package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums     []int
	target   int
	expected int
}

func TestGiven(t *testing.T) {
	tests := []Test{
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, search(test.nums, test.target))
	}
}
