package subtree_of_another_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	root    []int
	subRoot []int
	valid   bool
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

		if values[0] != NULL {
			current.Left = &TreeNode{Val: values[0]}
			q = append(q, current.Left)
		}
		values = values[1:]

		if len(values) > 0 {
			if values[0] != NULL {
				current.Right = &TreeNode{Val: values[0]}
				q = append(q, current.Right)
			}
			values = values[1:]
		}
	}

	return &root
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			root:    []int{3, 4, 5, 1, 2},
			subRoot: []int{4, 1, 2},
			valid:   true,
		},
		{
			root:    []int{3, 4, 5, 1, 2, NULL, NULL, NULL, NULL, 0},
			subRoot: []int{4, 1, 2},
			valid:   false,
		},
		{
			root:    []int{1, 1},
			subRoot: []int{1},
			valid:   true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.valid, isSubtree(toTree(test.root), toTree(test.subRoot)), "root: %v subroot: %v", test.root, test.subRoot)
	}
}
