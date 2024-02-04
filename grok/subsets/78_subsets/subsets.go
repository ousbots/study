package subsets

func subsets(nums []int) [][]int {
	return helper(nums, 0, []int{})
}

func helper(nums []int, start int, subset []int) [][]int {
	output := [][]int{subset}

	for i := start; i < len(nums); i++ {
		set := make([]int, len(subset))
		copy(set, subset)
		set = append(set, nums[i])
		output = append(output, helper(nums, i+1, set)...)
	}

	return output
}
