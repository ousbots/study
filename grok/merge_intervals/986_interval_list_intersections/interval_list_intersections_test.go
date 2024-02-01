package interval_list_intersections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	first  [][]int
	second [][]int
	valid  [][]int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			first:  [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			second: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			valid:  [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			first:  [][]int{{1, 3}, {5, 9}},
			second: [][]int{},
			valid:  [][]int{},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, intervalIntersection(test.first, test.second), "first: %v second: %v", test.first, test.second)
	}
}
