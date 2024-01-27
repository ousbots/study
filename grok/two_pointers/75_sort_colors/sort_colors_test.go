package sort_colors

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
			input: []int{2, 0, 2, 1, 1, 0},
			valid: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input: []int{2, 0, 1},
			valid: []int{0, 1, 2},
		},
		{
			input: []int{2, 2, 2, 1, 1, 0},
			valid: []int{0, 1, 1, 2, 2, 2},
		},
		{
			input: []int{2, 1, 0, 0, 1, 2},
			valid: []int{0, 0, 1, 1, 2, 2},
		},
	}

	for _, test := range tests {
		original := append([]int{}, test.input...)
		sortColors(test.input)
		assert.Equal(t, test.valid, test.input, "input: %v", original)
	}
}
