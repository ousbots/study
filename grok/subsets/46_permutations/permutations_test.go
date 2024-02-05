package permutations

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
			input: []int{1, 2, 3},
			valid: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			input: []int{0, 1},
			valid: [][]int{{0, 1}, {1, 0}},
		},
		{
			input: []int{1},
			valid: [][]int{{1}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, permute(test.input), "input: %v", test.input)
	}
}
