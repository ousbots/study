package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Determine if the given binary tree is height balanced, i.e. every subtree differs in height by
// no more than one.
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
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

		left := deepest(current.Left, 0)
		right := deepest(current.Right, 0)

		diff := left - right
		if diff < 0 {
			diff = -diff
		}

		if diff > 1 {
			return false
		}
	}

	return true
}

// Find the deepest node count of the given node.
func deepest(current *TreeNode, count int) int {
	if current == nil {
		return count
	}

	left := deepest(current.Left, count+1)
	right := deepest(current.Right, count+1)

	if left > right {
		return left
	} else {
		return right
	}
}
