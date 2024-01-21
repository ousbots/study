package symmetric_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	values   []int
	expected bool
}

const NULL = -99999

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

		current.Left = &TreeNode{Val: values[0]}
		values = values[1:]
		q = append(q, current.Left)

		if len(values) > 0 {
			current.Right = &TreeNode{Val: values[0]}
			values = values[1:]
			q = append(q, current.Right)
		}
	}

	return &root
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			values:   []int{1, 2, 2, 3, 4, 4, 3},
			expected: true,
		},
		{
			values:   []int{1, 2, 2, NULL, 3, NULL, 3},
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isSymmetric(toTree(test.values)), "values: %v", test.values)
	}
}
