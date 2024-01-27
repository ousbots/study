package linked_list_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input []int
	valid int
}

func toList(values []int, loop int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	root := ListNode{}
	current := &root
	var linkTo *ListNode

	for i, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
		if i == loop {
			linkTo = current
		}
	}

	if linkTo != nil {
		current.Next = linkTo
	}

	return root.Next
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: []int{3, 2, 0, -4},
			valid: 1,
		},
		{
			input: []int{1, 2},
			valid: 0,
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6},
			valid: 0,
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6},
			valid: 1,
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6},
			valid: 2,
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6, 7},
			valid: 0,
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6, 7},
			valid: 1,
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6, 7},
			valid: 2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.input[test.valid], detectCycle(toList(test.input, test.valid)).Val, "input: %v", test.input)
	}
}
