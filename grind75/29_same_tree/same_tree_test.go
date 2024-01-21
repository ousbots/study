package same_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	p        []int
	q        []int
	expected bool
}

const NULL = -9999

func toTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := TreeNode{Val: values[0]}
	values = values[1:]
	q := []*TreeNode{&root}

	for len(values) > 0 {
		cur := q[0]
		q = q[1:]

		cur.Left = &TreeNode{Val: values[0]}
		values = values[1:]
		q = append(q, cur.Left)

		if len(values) > 0 {
			cur.Right = &TreeNode{Val: values[0]}
			values = values[1:]
			q = append(q, cur.Right)
		}
	}

	return &root
}

func TestBasic(t *testing.T) {
	tests := []Test{
		{
			p:        []int{1, 2, 3},
			q:        []int{1, 2, 3},
			expected: true,
		},
		{
			p:        []int{1, 2},
			q:        []int{1, NULL, 2},
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isSameTree(toTree(test.p), toTree(test.q)))
	}
}
