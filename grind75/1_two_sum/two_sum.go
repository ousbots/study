package two_sum

// Find the indices of the two elements in the slice that add up to the target.
func twoSum(nums []int, target int) []int {
	values := map[int][]int{}

	for pos, val := range nums {
		values[val] = append(values[val], pos)
	}

	for i, val := range nums {
		positions, ok := values[target-val]
		if !ok {
			continue
		}

		for _, j := range positions {
			if i == j {
				continue
			}

			return []int{i, j}
		}
	}

	return []int{}
}

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
