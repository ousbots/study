package first_missing_positive

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
			input: []int{1, 2, 0},
			valid: 3,
		},
		{
			input: []int{3, 4, -1, 1},
			valid: 2,
		},
		{
			input: []int{7, 8, 9, 11, 12},
			valid: 1,
		},
		{
			input: []int{2, 3, 0, -1},
			valid: 1,
		},
		{
			input: []int{1},
			valid: 2,
		},
		{
			input: []int{1, 2, 3},
			valid: 4,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, firstMissingPositive(test.input), "input: %v", test.input)
	}
}
