package find_smallest_letter_greater_than_target

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input  []byte
	target byte
	valid  byte
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'a',
			valid:  'c',
		},
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'c',
			valid:  'f',
		},
		{
			input:  []byte{'x', 'x', 'y', 'y'},
			target: 'z',
			valid:  'x',
		},
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'j',
			valid:  'c',
		},
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'z',
			valid:  'c',
		},
		{
			input:  []byte{'e', 'e', 'g', 'g'},
			target: 'g',
			valid:  'e',
		},
		{
			input:  []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'},
			target: 'e',
			valid:  'n',
		},
	}

	for _, test := range tests {
		output := nextGreatestLetter(test.input, test.target)
		assert.Equal(t, test.valid, output, "input: %s, target: %c, valid: %c, output: %c", test.input, test.target, test.valid, output)
	}
}
