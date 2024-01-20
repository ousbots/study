package ransom_note

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	note     string
	magazine string
	expected bool
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			note:     "a",
			magazine: "b",
			expected: false,
		},
		{
			note:     "aa",
			magazine: "ab",
			expected: false,
		},
		{
			note:     "aa",
			magazine: "aab",
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, canConstruct(test.note, test.magazine))
	}
}
