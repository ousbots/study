package subsets_ii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(a, b int) bool { return nums[a] < nums[b] })
	return helper(nums, 0, []int{})
}

func helper(nums []int, start int, subset []int) [][]int {
	output := [][]int{subset}

	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}

		temp := make([]int, len(subset)+1)
		copy(temp, subset)
		temp[len(temp)-1] = nums[i]

		output = append(output, helper(nums, i+1, temp)...)
	}

	return output
}
