package add_binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	a        string
	b        string
	expected string
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, addBinary(test.a, test.b))
	}
}
