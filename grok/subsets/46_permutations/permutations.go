package permutations

func permute(nums []int) [][]int {
	return helper(nums, []int{})
}

func helper(nums []int, subset []int) [][]int {
	if len(nums) == 0 {
		return [][]int{subset}
	}

	output := [][]int{}

	for i := range nums {
		temp := make([]int, len(subset)+1)
		copy(temp, subset)
		temp[len(temp)-1] = nums[i]

		remainder := make([]int, len(nums))
		copy(remainder, nums)
		remainder = append(remainder[:i], remainder[i+1:]...)

		output = append(output, helper(remainder, temp)...)
	}

	return output
}
