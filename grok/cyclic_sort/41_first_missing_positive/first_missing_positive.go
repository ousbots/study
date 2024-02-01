package first_missing_positive

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i]-1 != i && nums[i] != nums[nums[i]-1] {
			temp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = temp
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
