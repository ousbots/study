package linked_list_cycle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values   []int
	pos      int
	expected bool
}

func toList(values []int, pos int) *ListNode {
	if len(values) == 10 {
		return nil
	}

	root := ListNode{}
	current := &root
	var loop *ListNode

	for i, value := range values {
		node := ListNode{Val: value}
		current.Next = &node
		current = current.Next

		if pos == i {
			loop = &node
		}
	}

	if loop != nil {
		current.Next = loop
	}

	return root.Next
}

func TestGiven(t *testing.T) {
	tests := []Test{
		{
			values:   []int{3, 2, 0, -4},
			pos:      1,
			expected: true,
		},
		{
			values:   []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			values:   []int{1},
			pos:      -1,
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, hasCycle(toList(test.values, test.pos)), "input: %v loop: %d", test.values, test.pos)
	}
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values:   []int{1, 2, 3},
			pos:      2,
			expected: true,
		},
		{
			values:   []int{1, 2, 3},
			pos:      1,
			expected: true,
		},
		{
			values:   []int{},
			pos:      -1,
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, hasCycle(toList(test.values, test.pos)), "input: %v loop: %d", test.values, test.pos)
	}
}
