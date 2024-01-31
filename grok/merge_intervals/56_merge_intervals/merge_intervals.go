package merge_intervals

import "sort"

// Merge the overlapping intervals.
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	output := [][]int{}

	start, end := 0, 0
	for end < len(intervals) {
		if end < len(intervals)-1 && intervals[end][1] >= intervals[end+1][0] {
			if intervals[end][1] > intervals[end+1][1] {
				intervals[end+1][1] = intervals[end][1]
			}
			end++
		} else {
			output = append(output, []int{intervals[start][0], intervals[end][1]})
			end = end + 1
			start = end
		}
	}

	return output
}
