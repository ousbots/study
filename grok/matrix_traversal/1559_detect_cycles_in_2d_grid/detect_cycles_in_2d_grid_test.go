package detect_cycles_in_2d_grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input [][]byte
	valid bool
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}},
			valid: true,
		},
		{
			input: [][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}},
			valid: true,
		},
		{
			input: [][]byte{{'a', 'b', 'b'}, {'b', 'z', 'b'}, {'b', 'b', 'a'}},
			valid: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, containsCycle(test.input), "input: %v", test.input)
	}
}
