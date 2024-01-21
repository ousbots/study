package palindrome_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values   []int
	expected bool
}

func toList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	root := ListNode{}
	last := &root
	for _, val := range values {
		node := ListNode{Val: val}
		last.Next = &node
		last = last.Next
	}

	return root.Next
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values:   []int{1, 2, 2, 1},
			expected: true,
		},
		{
			values:   []int{1, 2},
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isPalindrome(toList(test.values)), "values: %v", test.values)
	}
}
