package lowest_common_ancestor_of_a_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values   []int
	p        int
	q        int
	expected int
}

// Note: no re-balancing, values need to be in balanced order.
func toTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := TreeNode{Val: values[0]}
	values = values[1:]

	for _, val := range values {
		node := TreeNode{Val: val}

		look := &root
		for look != nil {
			if node.Val > look.Val {
				if look.Right == nil {
					look.Right = &node
					break
				} else {
					look = look.Right
				}
			} else {
				if look.Left == nil {
					look.Left = &node
					break
				} else {
					look = look.Left
				}
			}
		}
	}

	return &root
}

func TestGiven(t *testing.T) {
	tests := []Test{
		{
			values:   []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			p:        2,
			q:        8,
			expected: 6,
		},
		{
			values:   []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			p:        2,
			q:        4,
			expected: 2,
		},
		{
			values:   []int{2, 1},
			p:        2,
			q:        1,
			expected: 2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, lowestCommonAncestor(toTree(test.values), &TreeNode{Val: test.p}, &TreeNode{Val: test.q}).Val)
	}
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values:   []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			p:        0,
			q:        3,
			expected: 2,
		},
		{
			values:   []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			p:        0,
			q:        5,
			expected: 2,
		},
		{
			values:   []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			p:        0,
			q:        7,
			expected: 6,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, lowestCommonAncestor(toTree(test.values), &TreeNode{Val: test.p}, &TreeNode{Val: test.q}).Val)
	}
}
