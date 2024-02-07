package remove_duplicate_letters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input string
	valid string
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: "bcabc",
			valid: "abc",
		},
		{
			input: "cbacdcbc",
			valid: "acdb",
		},
		{
			input: "abacb",
			valid: "abc",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, removeDuplicateLetters(test.input), "input: %v", test.input)
	}
}
