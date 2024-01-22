package reverse_bits

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type Test struct {
	num      uint32
	expected uint32
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			num:      0b00000010100101000001111010011100,
			expected: 0b00111001011110000010100101000000,
		},
		{
			num:      0b11111111111111111111111111111101,
			expected: 0b10111111111111111111111111111111,
		},
	}

	for _, test := range tests {
		fmt.Printf("gen: %v\n", strconv.FormatInt(int64(reverseBits(test.num)), 2))
		fmt.Printf("exp: %v\n", strconv.FormatInt(int64(test.expected), 2))
		assert.Equal(t, test.expected, reverseBits(test.num))
	}
}
