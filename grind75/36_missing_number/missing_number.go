package missing_number

import "slices"

// Find the missing number from the sequence 0...n in the list.
func missingNumber(nums []int) int {
	total := (len(nums) * (len(nums) + 1)) / 2

	for _, val := range nums {
		total -= val
	}

	return total
}

func missingNumberSort(nums []int) int {
	slices.Sort(nums)

	for i, val := range nums {
		if i != val {
			return i
		}
	}

	return len(nums)
}
