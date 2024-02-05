package top_k_frequent_elements

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

type Test struct {
	input []int
	k     int
	valid []int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: []int{1, 1, 1, 2, 2, 3},
			k:     2,
			valid: []int{1, 2},
		},
		{
			input: []int{1},
			k:     1,
			valid: []int{1},
		},
		{
			input: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			k:     2,
			valid: []int{3, 4},
		},
		{
			input: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			k:     4,
			valid: []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		top := topKFrequent(test.input, test.k)
		slices.Sort(top)
		slices.Sort(test.valid)
		assert.Equal(t, test.valid, top, "input: %v k: %d", test.input, test.k)
	}
}
