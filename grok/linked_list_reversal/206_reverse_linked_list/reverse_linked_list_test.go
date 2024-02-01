package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input []int
	valid []int
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
			input: []int{1, 2, 3, 4, 5},
			valid: []int{5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, toSlice(reverseList(toList(test.input))), "input: %v", test.input)
	}
}
