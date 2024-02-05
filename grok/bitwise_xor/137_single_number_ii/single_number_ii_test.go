package single_number_ii

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
			input: []int{2, 2, 3, 2},
			valid: 3,
		},
		{
			input: []int{0, 1, 0, 1, 0, 1, 99},
			valid: 99,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, singleNumber(test.input), "input: %v", test.input)
	}
}
