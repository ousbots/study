package maximum_depth_of_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values   []int
	expected int
}

const MAGIC = -99999

func toTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := TreeNode{Val: values[0]}
	values = values[1:]
	q := []*TreeNode{&root}

	for len(values) > 0 {
		current := q[0]
		q = q[1:]

		val := values[0]
		values = values[1:]

		if val != MAGIC {
			current.Left = &TreeNode{Val: val}
			q = append(q, current.Left)
		}

		if len(values) > 0 {
			val = values[0]
			values = values[1:]

			if val != MAGIC {
				current.Right = &TreeNode{Val: val}
				q = append(q, current.Right)
			}
		}
	}

	return &root
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values:   []int{3, 9, 20, MAGIC, MAGIC, 15, 7},
			expected: 3,
		},
		{
			values:   []int{1, MAGIC, 2},
			expected: 2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, maxDepth(toTree(test.values)))
	}
}
