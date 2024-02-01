package convert_to_base_neg_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input int
	valid string
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: 2,
			valid: "110",
		},
		{
			input: 3,
			valid: "111",
		},
		{
			input: 4,
			valid: "100",
		},
		{
			input: 0,
			valid: "0",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, baseNeg2(test.input), "input: %v", test.input)
	}
}
