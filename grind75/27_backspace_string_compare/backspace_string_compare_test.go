package backspace_string_compare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	s        string
	t        string
	expected bool
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			s:        "ab#c",
			t:        "ad#c",
			expected: true,
		},
		{
			s:        "ab##",
			t:        "c#d#",
			expected: true,
		},
		{
			s:        "a#c",
			t:        "b",
			expected: false,
		},
		{
			s:        "ab#cdef###gh###",
			t:        "bc#def###q##ab#cde##fg###",
			expected: true,
		},
		{
			s:        "",
			t:        "#abc###",
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, backspaceCompare(test.s, test.t))
	}
}
