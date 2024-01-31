package longest_substring_with_at_least_k_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	s     string
	k     int
	valid int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			s:     "aaabb",
			k:     3,
			valid: 3,
		},
		{
			s:     "ababbc",
			k:     2,
			valid: 5,
		},
		{
			s:     "aaabbb",
			k:     3,
			valid: 6,
		},
		{
			s:     "aaaaaa",
			k:     6,
			valid: 6,
		},
		{
			s:     "aaaaaa",
			k:     5,
			valid: 6,
		},
		{
			s:     "baaaaaa",
			k:     5,
			valid: 6,
		},
		{
			s:     "aaaaaab",
			k:     5,
			valid: 6,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, longestSubstring(test.s, test.k), "s: %s k: %d", test.s, test.k)
	}
}
