package remove_all_adjacent_duplicates_in_string

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
			input: "abbaca",
			valid: "ca",
		},
		{
			input: "azxxzy",
			valid: "ay",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, removeDuplicates(test.input), "input: %v", test.input)
	}
}
