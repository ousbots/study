package next_greater_element_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	nums1 []int
	nums2 []int
	valid []int
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			valid: []int{-1, 3, -1},
		},
		{
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			valid: []int{3, -1},
		},
		{
			nums1: []int{1, 4, 2, 5},
			nums2: []int{5, 4, 3, 2, 1},
			valid: []int{-1, -1, -1, -1},
		},
		{
			nums1: []int{1, 4, 2, 5},
			nums2: []int{1, 2, 3, 4, 5},
			valid: []int{2, 5, 3, -1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, nextGreaterElement(test.nums1, test.nums2), "nums1: %v nums2: %v", test.nums1, test.nums2)
	}
}
