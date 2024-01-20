package first_bad_version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	max      int
	bad      int
	expected int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			max:      5,
			bad:      4,
			expected: 4,
		},
		{
			max:      9,
			bad:      9,
			expected: 9,
		},
		{
			max:      9,
			bad:      1,
			expected: 1,
		},
		{
			max:      9,
			bad:      2,
			expected: 2,
		},
	}

	for _, test := range tests {
		badVersion = test.bad
		assert.Equal(t, test.expected, firstBadVersion(test.max), "max version %d", test.max)
	}
}
