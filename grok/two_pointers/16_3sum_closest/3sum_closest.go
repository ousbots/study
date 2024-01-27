package _3sum_closest

import "slices"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		panic("not enough numbers")
	}

	slices.Sort(nums)
	closest := nums[0] + nums[1] + nums[2]

	for left := 0; left <= len(nums)-3; left++ {
		middle := left + 1
		right := len(nums) - 1

		for middle < right {
			temp := nums[left] + nums[middle] + nums[right]

			if temp == target {
				return temp
			}

			if abs(target-temp) < abs(target-closest) {
				closest = temp
			}

			if temp > target {
				right -= 1
			} else {
				middle += 1
			}
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func diff(a, b int) int {
	if a-b < 0 {
		return b - a
	}

	return a - b
}

func threeSumClosestBruteForce(nums []int, target int) int {
	closest := -9999999
	closest_diff := 9999999

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]

				diff := target - sum
				if diff < 0 {
					diff = -diff
				}

				if diff < closest_diff {
					closest = sum
					closest_diff = diff
				}
			}
		}
	}

	return closest
}
