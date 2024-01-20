package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Find the longest path between any two nodes in a binary tree.
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return -1
	}

	diameter := -1
	q := []*TreeNode{root}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		diam := depth(current.Left, 0) + depth(current.Right, 0)
		if diam > diameter {
			diameter = diam
		}

		if current.Left != nil {
			q = append(q, current.Left)
		}

		if current.Right != nil {
			q = append(q, current.Right)
		}
	}

	return diameter
}

func depth(current *TreeNode, level int) int {
	if current == nil {
		return level
	}

	left := depth(current.Left, level+1)
	right := depth(current.Right, level+1)

	if left > right {
		return left
	} else {
		return right
	}
}
