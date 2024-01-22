package squares_of_a_sorted_array

import "slices"

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	last := len(nums) - 1
	for nums[0] < 0 {
		temp := -nums[0]
		for last > 0 && temp < nums[last] {
			last--
		}

		nums = append(nums[1:last+1], nums[last:]...)
		nums[last] = temp
	}

	for i := range nums {
		nums[i] *= nums[i]
	}

	return nums
}

func sortedSquaresSimple(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}

	slices.Sort(nums)
	return nums
}
