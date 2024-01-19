package flood_fill

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	image    [][]int
	sr       int
	sc       int
	color    int
	expected [][]int
}

func TestGiven(t *testing.T) {
	tests := []Test{
		{
			image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:       1,
			sc:       1,
			color:    2,
			expected: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			image:    [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:       0,
			sc:       0,
			color:    0,
			expected: [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, floodFill(test.image, test.sr, test.sc, test.color))
	}
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			image:    [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}},
			sr:       1,
			sc:       1,
			color:    2,
			expected: [][]int{{1, 0, 1}, {0, 2, 0}, {1, 0, 1}},
		},
		{
			image:    [][]int{{1, 0, 1, 1}, {1, 1, 0, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}},
			sr:       1,
			sc:       1,
			color:    2,
			expected: [][]int{{2, 0, 2, 2}, {2, 2, 0, 2}, {2, 0, 2, 2}, {2, 2, 2, 2}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, floodFill(test.image, test.sr, test.sc, test.color))
	}
}
