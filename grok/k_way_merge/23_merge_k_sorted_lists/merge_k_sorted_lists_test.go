package merge_k_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input [][]int
	valid []int
}

func toList(vals [][]int) []*ListNode {
	output := []*ListNode{}
	for _, list := range vals {
		root := ListNode{}
		current := &root

		for _, val := range list {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}

		output = append(output, root.Next)
	}

	return output
}

func toSlice(root *ListNode) []int {
	output := []int{}

	for root != nil {
		output = append(output, root.Val)
		root = root.Next
	}

	return output
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			valid: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			input: [][]int{},
			valid: []int{},
		},
		{
			input: [][]int{{}},
			valid: []int{},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, toSlice(mergeKLists(toList(test.input))), "input: %v", test.input)
	}
}
