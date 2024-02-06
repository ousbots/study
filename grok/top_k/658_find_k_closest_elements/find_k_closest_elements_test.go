package find_k_closest_elements

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

type Test struct {
	arr   []int
	k     int
	x     int
	valid []int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			arr:   []int{1, 2, 3, 4, 5},
			k:     4,
			x:     3,
			valid: []int{1, 2, 3, 4},
		},
		{
			arr:   []int{1, 2, 3, 4, 5},
			k:     4,
			x:     -1,
			valid: []int{1, 2, 3, 4},
		},
		{
			arr:   []int{1, 2, 3, 4, 5},
			k:     2,
			x:     -1,
			valid: []int{1, 2},
		},
		{
			arr:   []int{1, 2, 3, 4, 5},
			k:     2,
			x:     6,
			valid: []int{4, 5},
		},
		{
			arr:   []int{-2, -1, 0, 1, 2},
			k:     3,
			x:     0,
			valid: []int{-1, 0, 1},
		},
		{
			arr:   []int{-2, -1, 0, 1, 2},
			k:     4,
			x:     0,
			valid: []int{-2, -1, 0, 1},
		},
	}

	for _, test := range tests {
		slices.Sort(test.valid)
		assert.Equal(t, test.valid, findClosestElements(test.arr, test.k, test.x), "arr: %v k: %d x: %d", test.arr, test.k, test.x)
	}
}
