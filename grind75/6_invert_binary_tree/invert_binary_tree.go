package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Invert a binary tree.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		current := q[0]
		q = q[1:]

		if current.Left != nil {
			q = append(q, current.Left)
		}

		if current.Right != nil {
			q = append(q, current.Right)
		}

		temp := current.Left
		current.Left = current.Right
		current.Right = temp
	}

	return root
}
