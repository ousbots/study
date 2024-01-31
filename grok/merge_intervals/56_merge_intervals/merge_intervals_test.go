package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input [][]int
	valid [][]int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			valid: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input: [][]int{{1, 4}, {4, 5}},
			valid: [][]int{{1, 5}},
		},
		{
			input: [][]int{{1, 8}, {2, 3}, {4, 5}, {10, 11}},
			valid: [][]int{{1, 8}, {10, 11}},
		},
		{
			input: [][]int{{1, 5}, {2, 3}, {4, 6}, {10, 11}},
			valid: [][]int{{1, 6}, {10, 11}},
		},
		{
			input: [][]int{{1, 20}, {2, 3}, {4, 6}, {10, 11}},
			valid: [][]int{{1, 20}},
		},
		{
			input: [][]int{{2, 3}, {1, 9}, {4, 12}, {10, 11}},
			valid: [][]int{{1, 12}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, merge(test.input), "input: %v", test.input)
	}
}
