package minimum_add_to_make_parentheses_valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input string
	valid int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: "())",
			valid: 1,
		},
		{
			input: "(((",
			valid: 3,
		},
		{
			input: "))((",
			valid: 4,
		},
		{
			input: "(()))((",
			valid: 3,
		},
		{
			input: "()()(())(",
			valid: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, minAddToMakeValid(test.input), "input: %v", test.input)
	}
}
