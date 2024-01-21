package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Determine if a binary tree is left-right symmetric.
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftq := []*TreeNode{root.Left}
	rightq := []*TreeNode{root.Right}

	for len(leftq) > 0 && len(rightq) > 0 {
		left := leftq[0]
		right := rightq[0]

		leftq = leftq[1:]
		rightq = rightq[1:]

		if left == nil || right == nil {
			if left == nil && right == nil {
				continue
			}

			return false
		}

		if left.Val != right.Val {
			return false
		}

		leftq = append(leftq, left.Left)
		leftq = append(leftq, left.Right)
		rightq = append(rightq, right.Right)
		rightq = append(rightq, right.Left)
	}

	if len(leftq) > 0 || len(rightq) > 0 {
		return false
	}

	return true
}
