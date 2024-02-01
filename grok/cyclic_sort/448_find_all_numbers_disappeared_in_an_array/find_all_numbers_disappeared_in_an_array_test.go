package find_all_numbers_disappeared_in_an_array

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
			input: []int{4, 3, 2, 7, 8, 2, 3, 1},
			valid: []int{5, 6},
		},
		{
			input: []int{1, 1},
			valid: []int{2},
		},
		{
			input: []int{3, 2, 3},
			valid: []int{1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, findDisappearedNumbers(test.input), "input: %v", test.input)
	}
}
