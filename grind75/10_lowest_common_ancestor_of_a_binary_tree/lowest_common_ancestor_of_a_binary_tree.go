package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	current := root

	for current != nil {
		if p.Val < current.Val && q.Val < current.Val {
			current = current.Left
			continue
		}

		if p.Val > current.Val && q.Val > current.Val {
			current = current.Right
			continue
		}

		return current
	}

	return nil
}
