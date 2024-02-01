package simplify_path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input string
	valid string
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: "/home/",
			valid: "/home",
		},
		{
			input: "/../",
			valid: "/",
		},
		{
			input: "/home//foo/",
			valid: "/home/foo",
		},
		{
			input: "/.././../",
			valid: "/",
		},
		{
			input: "/what/./is/this/stinky/../doodoo",
			valid: "/what/is/this/doodoo",
		},
		{
			input: "/what/./is/this/stinky/../../doodoo",
			valid: "/what/is/doodoo",
		},
		{
			input: "this/is/not/really/../../../../real",
			valid: "/real",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, simplifyPath(test.input), "input: %v", test.input)
	}
}
