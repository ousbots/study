package majority_element

func majorityElement(nums []int) int {
	count := map[int]int{}

	for _, num := range nums {
		count[num]++
	}

	max := 0
	element := 0

	for elem, total := range count {
		if total > max {
			max = total
			element = elem
		}
	}

	return element
}
