package interval_list_intersections

// Generate the intersection of two intervals.
func intervalIntersection(firstList, secondList [][]int) [][]int {
	output := [][]int{}
	first, second := 0, 0

	for first < len(firstList) && second < len(secondList) {
		if firstList[first][1] < secondList[second][0] {
			first++
		} else if secondList[second][1] < firstList[first][0] {
			second++
		} else {
			var start, end int

			if firstList[first][0] >= secondList[second][0] {
				start = firstList[first][0]
			} else {
				start = secondList[second][0]
			}

			if firstList[first][1] <= secondList[second][1] {
				end = firstList[first][1]
			} else {
				end = secondList[second][1]
			}

			if firstList[first][1] < secondList[second][1] {
				first++
			} else {
				second++
			}

			output = append(output, []int{start, end})
		}
	}

	return output
}
