package _3sum_closest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums   []int
	target int
	valid  int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			valid:  2,
		},
		{
			nums:   []int{0, 0, 0},
			target: 0,
			valid:  0,
		},
		{
			nums:   []int{0, 1, 2},
			target: 0,
			valid:  3,
		},
		{
			nums:   []int{0, 3, 97, 102, 200},
			target: 300,
			valid:  300,
		},
		{
			nums:   []int{0, 4, 97, 101, 200},
			target: 300,
			valid:  301,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, threeSumClosest(test.nums, test.target), "nums: %v target: %d", test.nums, test.target)
	}
}
