package squares_of_a_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums  []int
	valid []int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			nums:  []int{-4, -1, 0, 3, 10},
			valid: []int{0, 1, 9, 16, 100},
		},
		{
			nums:  []int{-7, -3, 2, 3, 11},
			valid: []int{4, 9, 9, 49, 121},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, sortedSquares(test.nums), "input: %v", test.nums)
	}
}
