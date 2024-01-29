package happy_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input int
	valid bool
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: 19,
			valid: true,
		},
		{
			input: 2,
			valid: false,
		},
		{
			input: 1,
			valid: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, isHappy(test.input), "input: %v", test.input)
	}
}
