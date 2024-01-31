package fruit_into_baskets

func totalFruit(fruits []int) int {
	max := 0
	count := make([]int, len(fruits))
	numFruit := 0
	start, end := 0, 0

	for end < len(fruits) {
		if numFruit > 2 {
			count[fruits[start]]--
			if count[fruits[start]] == 0 {
				numFruit--
			}
			start++
		} else {
			count[fruits[end]]++
			if count[fruits[end]] == 1 {
				numFruit++
			}
			end++
		}

		if numFruit <= 2 {
			if end-start > max {
				max = end - start
			}
		}
	}

	return max
}
