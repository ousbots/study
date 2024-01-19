package balanced_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values   []int
	expected bool
}

const MAGIC = -9999999999

func toTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := TreeNode{Val: values[0]}
	values = values[1:]

	q := []*TreeNode{&root}

	for len(q) != 0 {
		current := q[0]
		q = q[1:]

		if len(values) > 0 {
			val := values[0]
			values = values[1:]

			if val != MAGIC {
				current.Left = &TreeNode{Val: val}
				q = append(q, current.Left)
			}
		}

		if len(values) > 0 {
			val := values[0]
			values = values[1:]

			if val != MAGIC {
				current.Right = &TreeNode{Val: val}
				q = append(q, current.Right)
			}
		}

		if len(values) == 0 {
			break
		}
	}

	return &root
}

func TestGiven(t *testing.T) {
	tests := []Test{
		{
			values:   []int{3, 9, 20, MAGIC, MAGIC, 15, 7},
			expected: true,
		},
		{
			values:   []int{1, 2, 2, 3, 3, MAGIC, MAGIC, 4, 4},
			expected: false,
		},
		{
			values:   []int{},
			expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isBalanced(toTree(test.values)))
	}
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values:   []int{1, 2, 2, 3, MAGIC, MAGIC, 3, 4, MAGIC, MAGIC, 4},
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isBalanced(toTree(test.values)))
	}
}
