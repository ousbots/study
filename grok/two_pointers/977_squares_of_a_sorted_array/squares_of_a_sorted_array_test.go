package squares_of_a_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input []int
	valid []int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: []int{},
			valid: []int{},
		},
		{
			input: []int{-4, -1, 0, 3, 10},
			valid: []int{0, 1, 9, 16, 100},
		},
		{
			input: []int{-7, -3, 2, 3, 11},
			valid: []int{4, 9, 9, 49, 121},
		},
		{
			input: []int{-7, -3, 6},
			valid: []int{9, 36, 49},
		},
		{
			input: []int{-7, -3, 7},
			valid: []int{9, 49, 49},
		},
		{
			input: []int{2},
			valid: []int{4},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, sortedSquares(test.input), "input: %v", test.input)
	}
}
