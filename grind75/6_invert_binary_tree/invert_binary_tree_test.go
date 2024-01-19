package invert_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input    []int
	expected []int
}

func toSlice(root *TreeNode) []int {
	q := []*TreeNode{root}
	values := []int{}

	for len(q) != 0 {
		node := q[0]
		q = q[1:]

		if node != nil {
			q = append(q, node.Left)
			q = append(q, node.Right)
			values = append(values, node.Val)
		} else {
			values = append(values, -999)
		}
	}

	for i := len(values) - 1; i >= 0; i-- {
		if values[i] == -999 {
			values = values[:i]
		} else {
			break
		}
	}

	return values
}

func toTree(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	root := TreeNode{Val: nodes[0]}
	nodes = nodes[1:]

	q := []*TreeNode{&root}
	for len(nodes) != 0 {
		current := q[0]
		q = q[1:]

		if len(nodes) != 0 {
			left := nodes[0]
			nodes = nodes[1:]

			if left != -999 {
				current.Left = &TreeNode{Val: left}
				q = append(q, current.Left)
			}
		}

		if len(nodes) != 0 {
			right := nodes[0]
			nodes = nodes[1:]

			if right != -999 {
				current.Right = &TreeNode{Val: right}
				q = append(q, current.Right)
			}
		}
	}

	return &root
}

func TestGiven(t *testing.T) {
	tests := []Test{
		{
			input:    []int{4, 2, 7, 1, 3, 6, 9},
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			input:    []int{2, 1, 3},
			expected: []int{2, 3, 1},
		},
		{
			input:    []int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, toSlice(invertTree(toTree(test.input))))
	}
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			input:    []int{1, -999, 2, 3, -999, -999, 4},
			expected: []int{1, 2, -999, -999, 3, 4},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, toSlice(invertTree(toTree(test.input))))
	}
}
