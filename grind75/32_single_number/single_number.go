package single_number

// Return the number that only occurs once in the list.
func singleNumber(nums []int) int {
	seen := map[int]int{}

	for _, val := range nums {
		seen[val]++
	}

	for val, count := range seen {
		if count == 1 {
			return val
		}
	}

	return -1
}
