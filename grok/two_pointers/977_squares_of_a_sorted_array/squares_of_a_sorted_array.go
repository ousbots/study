package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	first := 0
	last := len(nums) - 1
	output := make([]int, len(nums))
	i := len(output) - 1

	for first <= last {
		if nums[first] < 0 {
			if -nums[first] > nums[last] {
				output[i] = nums[first] * nums[first]
				first++
			} else {
				output[i] = nums[last] * nums[last]
				last--
			}
		} else {
			output[i] = nums[last] * nums[last]
			last--
		}

		i--
	}

	return output
}
