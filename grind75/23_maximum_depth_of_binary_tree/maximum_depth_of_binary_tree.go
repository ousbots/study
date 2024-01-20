package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return depth(root, 0)
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
