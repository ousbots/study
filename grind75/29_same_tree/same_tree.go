package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Determine if the two trees are the same.
func isSameTree(s, t *TreeNode) bool {
	sq := []*TreeNode{s}
	tq := []*TreeNode{t}

	for len(sq) > 0 && len(tq) > 0 {
		scur := sq[0]
		tcur := tq[0]

		sq = sq[1:]
		tq = tq[1:]

		if scur == nil || tcur == nil {
			if scur == nil && tcur == nil {
				continue
			} else {
				return false
			}
		}

		if scur.Val != tcur.Val {
			return false
		}

		sq = append(sq, scur.Left, scur.Right)
		tq = append(tq, tcur.Left, tcur.Right)
	}

	return true
}
