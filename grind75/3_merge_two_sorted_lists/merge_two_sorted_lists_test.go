package merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	list1    []int
	list2    []int
	expected []int
}

func toList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := []ListNode{}

	for _, value := range values {
		nodes = append(nodes, ListNode{Val: value, Next: nil})
	}

	for i := range nodes {
		if i < len(nodes)-1 {
			nodes[i].Next = &nodes[i+1]
		}
	}

	return &nodes[0]
}

func toSlice(root *ListNode) []int {
	values := []int{}

	for root != nil {
		values = append(values, root.Val)
		root = root.Next
	}

	return values
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
	}

	for _, test := range tests {
		merged := toSlice(mergeTwoLists(toList(test.list1), toList(test.list2)))
		assert.Equal(t, test.expected, merged)
	}
}

func TestMore(t *testing.T) {
	tests := []Test{
		{
			list1:    []int{1, 2, 3},
			list2:    []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			list1:    []int{1},
			list2:    []int{2, 3},
			expected: []int{1, 2, 3},
		},
		{
			list1:    []int{2, 5},
			list2:    []int{1, 3, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range tests {
		merged := toSlice(mergeTwoLists(toList(test.list1), toList(test.list2)))
		assert.Equal(t, test.expected, merged)
	}
}
