package reverse_linked_list_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values []int
	left   int
	right  int
	valid  []int
}

func toList(vals []int) *ListNode {
	root := ListNode{}
	current := &root

	for _, val := range vals {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return root.Next
}

func toSlice(root *ListNode) []int {
	vals := []int{}
	current := root

	for current != nil {
		vals = append(vals, current.Val)
		current = current.Next
	}

	return vals
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values: []int{1, 2, 3, 4, 5},
			left:   2,
			right:  4,
			valid:  []int{1, 4, 3, 2, 5},
		},
		{
			values: []int{5},
			left:   1,
			right:  1,
			valid:  []int{5},
		},
		{
			values: []int{1, 2, 3},
			left:   1,
			right:  2,
			valid:  []int{2, 1, 3},
		},
		{
			values: []int{1, 2, 3},
			left:   2,
			right:  3,
			valid:  []int{1, 3, 2},
		},
		{
			values: []int{1, 2, 3},
			left:   1,
			right:  3,
			valid:  []int{3, 2, 1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, toSlice(reverseBetween(toList(test.values), test.left, test.right)), "values: %v right: %d left: %d", test.values, test.left, test.right)
	}
}
