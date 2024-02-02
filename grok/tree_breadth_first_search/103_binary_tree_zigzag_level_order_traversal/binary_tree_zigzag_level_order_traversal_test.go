package binary_tree_zigzag_level_order_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input []int
	valid [][]int
}

const NULL = -99999

func toTree(values []int) *TreeNode {
	q := []*TreeNode{}
	var root *TreeNode

	for len(values) > 0 {
		if len(q) == 0 {
			root = &TreeNode{Val: values[0]}
			values = values[1:]
			q = append(q, root)
			continue
		}

		current := q[0]
		q = q[1:]

		val := values[0]
		values = values[1:]

		if val != NULL {
			current.Left = &TreeNode{Val: val}
			q = append(q, current.Left)
		}

		if len(values) > 0 {
			val := values[0]
			values = values[1:]

			if val != NULL {
				current.Right = &TreeNode{Val: val}
				q = append(q, current.Right)
			}
		}
	}

	return root
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input: []int{3, 9, 20, NULL, NULL, 15, 7},
			valid: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			input: []int{1},
			valid: [][]int{{1}},
		},
		{
			input: []int{},
			valid: [][]int{},
		},
		{
			input: []int{1, 2, 3, 4, NULL, NULL, 5},
			valid: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, zigzagLevelOrder(toTree(test.input)), "input: %v", test.input)
	}
}
