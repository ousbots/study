package number_of_islands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input [][]byte
	valid int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
			valid: 1,
		},
		{
			input: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
			valid: 3,
		},
		{
			input: [][]byte{{'1', '0', '1', '0', '1'}, {'0', '1', '0', '1', '0'}, {'1', '0', '1', '0', '1'}, {'0', '1', '0', '1', '0'}},
			valid: 10,
		},
		{
			input: [][]byte{{'1', '1', '1', '1', '1'}, {'0', '0', '0', '0', '1'}, {'0', '1', '1', '0', '1'}, {'0', '1', '0', '0', '1'}, {'0', '1', '1', '1', '1'}},
			valid: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, numIslands(test.input), "input: %v", test.input)
	}
}
