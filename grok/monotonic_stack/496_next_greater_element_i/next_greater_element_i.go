package next_greater_element_i

func nextGreaterElement(nums1, nums2 []int) []int {
	next := map[int]int{}
	greater := []int{}

	for _, val := range nums2 {
		for len(greater) > 0 && val > greater[len(greater)-1] {
			next[greater[len(greater)-1]] = val
			greater = greater[:len(greater)-1]
		}

		greater = append(greater, val)
	}

	for _, val := range greater {
		next[val] = -1
	}

	output := []int{}
	for _, check := range nums1 {
		output = append(output, next[check])
	}

	return output
}

func nextGreaterElementBrute(nums1, nums2 []int) []int {
	output := []int{}

	for _, target := range nums1 {
		next := -1
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums2[j] == target {
				break
			}

			if nums2[j] > target {
				next = nums2[j]
			}
		}

		output = append(output, next)
	}

	return output
}
