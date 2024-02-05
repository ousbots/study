package minimum_absolute_difference

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input []int
	valid [][]int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: []int{4, 2, 1, 3},
			valid: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			input: []int{1, 3, 6, 10, 15},
			valid: [][]int{{1, 3}},
		},
		{
			input: []int{3, 8, -10, 23, 19, -4, -14, 27},
			valid: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
		{
			input: []int{-3, -2, -1, 0, 1, 2, 3},
			valid: [][]int{{-3, -2}, {-2, -1}, {-1, 0}, {0, 1}, {1, 2}, {2, 3}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, minimumAbsDifference(test.input), "input: %v", test.input)
	}
}
