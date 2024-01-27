package sort_colors

func sortColors(nums []int) {
	if len(nums) == 0 {
		panic("array too small")
	}

	left := 0
	right := len(nums) - 1

	for i := range nums {
		if nums[i] == 0 {
			left++
		}

		if nums[i] == 2 {
			right--
		}
	}

	for i := range nums {
		if i < left {
			nums[i] = 0
		} else if i > right {
			nums[i] = 2
		} else {
			nums[i] = 1
		}
	}
}
