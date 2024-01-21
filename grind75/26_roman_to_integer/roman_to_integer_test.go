package roman_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input    string
	expected int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input:    "III",
			expected: 3,
		},
		{
			input:    "LVIII",
			expected: 58,
		},
		{
			input:    "MCMXCIV",
			expected: 1994,
		},
		{
			input:    "MMMMCDXLIV",
			expected: 4444,
		},
		{
			input:    "MMMMMMMMMCMXCIX",
			expected: 9999,
		},
		{
			input:    "LV",
			expected: 55,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, romanToInt(test.input), "input: %s", test.input)
	}
}
