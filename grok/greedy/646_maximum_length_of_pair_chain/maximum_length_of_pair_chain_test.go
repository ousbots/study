package maximum_length_of_pair_chain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input [][]int
	valid int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: [][]int{{1, 2}, {2, 3}, {3, 4}},
			valid: 2,
		},
		{
			input: [][]int{{1, 2}, {7, 8}, {4, 5}},
			valid: 3,
		},
		{
			input: [][]int{{1, 99}, {7, 8}, {4, 5}, {2, 3}},
			valid: 3,
		},
		{
			input: [][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}},
			valid: 4,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, findLongestChain(test.input), "input: %v", test.input)
	}
}
