package find_all_numbers_disappeared_in_an_array

// Given an array of integers [1, n], find the integers in the range that are missing.
func findDisappearedNumbers(nums []int) []int {
	missing := []int{}

	for i := range nums {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			temp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = temp
		}
	}

	for i, val := range nums {
		if i+1 != val {
			missing = append(missing, i+1)
		}
	}

	return missing
}
