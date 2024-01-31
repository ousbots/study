package fruit_into_baskets

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
			input: []int{1, 2, 1},
			valid: 3,
		},
		{
			input: []int{0, 1, 2, 2},
			valid: 3,
		},
		{
			input: []int{1, 2, 3, 2, 2},
			valid: 4,
		},
		{
			input: []int{1, 1, 2, 3, 3, 3},
			valid: 4,
		},
		{
			input: []int{1, 1, 1, 2, 3, 3},
			valid: 4,
		},
		{
			input: []int{0},
			valid: 1,
		},
		{
			input: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			valid: 7,
		},
		{
			input: []int{1, 1, 2, 2, 2, 3, 3, 3, 4, 4},
			valid: 6,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, totalFruit(test.input), "input: %v", test.input)
	}
}
