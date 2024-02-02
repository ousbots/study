package remove_all_adjacent_duplicates_in_string_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	s     string
	k     int
	valid string
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			s:     "abcd",
			k:     2,
			valid: "abcd",
		},
		{
			s:     "deeedbbcccbdaa",
			k:     3,
			valid: "aa",
		},
		{
			s:     "pbbcggttciiippooaais",
			k:     2,
			valid: "ps",
		},
		{
			s:     "abccddeeedcbbaaq",
			k:     3,
			valid: "q",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, removeDuplicates(test.s, test.k), "s: %s k: %d", test.s, test.k)
	}
}
