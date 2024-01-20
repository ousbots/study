package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values   []int
	expected []int
}

func toList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := []*ListNode{}

	for _, val := range values {
		nodes = append(nodes, &ListNode{Val: val})
	}

	for i := range nodes {
		if i < len(nodes)-1 {
			nodes[i].Next = nodes[i+1]
		}
	}

	return nodes[0]
}

func toSlice(head *ListNode) []int {
	values := []int{}

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	return values
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values:   []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			values:   []int{1, 2},
			expected: []int{2, 1},
		},
		{
			values:   []int{1},
			expected: []int{1},
		},
		{
			values:   []int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, toSlice(reverseList(toList(test.values))))
	}
}
