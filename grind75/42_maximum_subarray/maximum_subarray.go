package maximum_subarray

// Determine the maximum sum of any subarray.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	total := nums[0]
	start := 0
	end := 0

	for start < len(nums)-1 {
		if start == end {
			end++
			total += nums[end]
		} else {
			if nums[start] <= 0 || total <= 0 {
				total -= nums[start]
				start++
			} else {
				if end < len(nums)-1 {
					end++
					total += nums[end]
				} else {
					break
				}
			}
		}

		if total > max {
			max = total
		}
	}

	return max
}
