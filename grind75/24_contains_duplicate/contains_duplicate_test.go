package contains_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums     []int
	expected bool
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, containsDuplicate(test.nums))
	}
}
