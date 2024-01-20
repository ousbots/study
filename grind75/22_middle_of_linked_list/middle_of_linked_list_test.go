package middle_of_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values   []int
	expected int
}

func toList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	root := ListNode{Val: values[0]}
	values = values[1:]

	last := &root
	for _, val := range values {
		node := &ListNode{Val: val}
		last.Next = node
		last = node
	}

	return &root
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values:   []int{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			values:   []int{1, 2, 3, 4, 5, 6},
			expected: 4,
		},
		{
			values:   []int{1},
			expected: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, middleNode(toList(test.values)).Val, "values: %v", test.values)
	}
}
