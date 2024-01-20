package majority_element

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
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, majorityElement(test.nums))
	}
}
