package contains_duplicate

import "slices"

// Determine if a list of numbers contains a duplicate.
func containsDuplicate(nums []int) bool {
	slices.Sort(nums)

	for i := range nums {
		if i > 0 {
			if nums[i] == nums[i-1] {
				return true
			}
		}
	}

	return false
}
