package subtree_of_another_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Determine if the subtree is found in the tree.
// Note: this is slow.
func isSubtree(root, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}

	q := []*TreeNode{root}
loop:
	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		if current == nil {
			continue
		}

		q = append(q, current.Left, current.Right)

		if current.Val == subRoot.Val {
			checkq := []*TreeNode{current}
			subq := []*TreeNode{subRoot}

			for len(checkq) > 0 && len(subq) > 0 {
				check := checkq[0]
				sub := subq[0]

				checkq = checkq[1:]
				subq = subq[1:]

				if check == nil || sub == nil {
					if check != sub {
						continue loop
					}
					continue
				}

				if check.Val != sub.Val {
					continue loop
				}

				checkq = append(checkq, check.Left, check.Right)
				subq = append(subq, sub.Left, sub.Right)
			}

			if len(checkq) == 0 && len(subq) == 0 {
				return true
			}
		}
	}

	return false
}
