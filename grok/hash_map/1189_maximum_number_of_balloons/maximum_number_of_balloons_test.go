package maximum_number_of_balloons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input string
	valid int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: "nlaebolko",
			valid: 1,
		},
		{
			input: "loonbalxballpoon",
			valid: 2,
		},
		{
			input: "leetcode",
			valid: 0,
		},
		{
			input: "a",
			valid: 0,
		},
		{
			input: "lloo",
			valid: 0,
		},
		{
			input: "balloonballoonballoonballoonballoonballooballooballooballooballooballoo",
			valid: 5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, maxNumberOfBalloons(test.input), "input: %v", test.input)
	}
}
