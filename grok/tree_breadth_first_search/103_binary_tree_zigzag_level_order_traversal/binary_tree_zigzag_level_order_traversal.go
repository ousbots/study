package binary_tree_zigzag_level_order_traversal

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	current := []*TreeNode{root}
	children := []*TreeNode{}
	output := [][]int{}
	level := []int{}
	reverse := false

	for len(current) != 0 || len(children) != 0 {
		if len(current) == 0 {
			current = children
			children = []*TreeNode{}
			output = append(output, level)
			slices.Reverse(current)
			level = []int{}
			reverse = !reverse
		}

		temp := current[0]
		current = current[1:]

		if temp == nil {
			continue
		}

		level = append(level, temp.Val)

		if reverse {
			children = append(children, temp.Right)
			children = append(children, temp.Left)
		} else {
			children = append(children, temp.Left)
			children = append(children, temp.Right)
		}
	}

	return output
}
