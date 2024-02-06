package kth_smallest_element_in_a_sorted_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input [][]int
	k     int
	valid int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}},
			k:     8,
			valid: 13,
		},
		{
			input: [][]int{{-5}},
			k:     1,
			valid: -5,
		},
		{
			input: [][]int{{1, 2}, {3, 3}},
			k:     1,
			valid: 1,
		},
		{
			input: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			k:     20,
			valid: 21,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, kthSmallest(test.input, test.k), "input: %v k: %d", test.input, test.k)
	}
}
