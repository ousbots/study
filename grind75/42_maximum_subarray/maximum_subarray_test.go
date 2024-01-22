package maximum_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input []int
	valid int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			valid: 6,
		},
		{
			input: []int{1},
			valid: 1,
		},
		{
			input: []int{5, 4, -1, 7, 8},
			valid: 23,
		},
		{
			input: []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4},
			valid: 6,
		},
		{
			input: []int{-1, 2, 2, -3},
			valid: 4,
		},
		{
			input: []int{9, -1, -1, -1},
			valid: 9,
		},
		{
			input: []int{-1, -1, -1, 9},
			valid: 9,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, maxSubArray(test.input), "input: %v", test.input)
	}
}
