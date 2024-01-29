package palindrome_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input []int
	valid bool
}

func toList(values []int) *ListNode {
	root := ListNode{}
	current := &root

	for _, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return root.Next
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: []int{1, 2, 2, 1},
			valid: true,
		},
		{
			input: []int{1, 2},
			valid: false,
		},
		{
			input: []int{1, 2, 1},
			valid: true,
		},
		{
			input: []int{2},
			valid: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, isPalindrome(toList(test.input)), "input: %v", test.input)
	}
}
