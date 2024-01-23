package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	start := 0
	end := 0

	for start < len(intervals)-1 {
		if newInterval[0] <= intervals[start][1] {
			break
		} else {
			start++
		}
	}

	for end < len(intervals)-1 {
		if newInterval[1] < intervals[end+1][0] {
			break
		}
		end++
	}

	if end <= start {
		if newInterval[1] < intervals[start][0] {
			intervals = append(intervals[:start+1], intervals[start:]...)
			intervals[start] = newInterval
		} else if newInterval[0] > intervals[start][1] {
			intervals = append(intervals[:start+1], intervals[start:]...)
			intervals[start+1] = newInterval
		} else {
			if newInterval[0] < intervals[start][0] {
				intervals[start][0] = newInterval[0]
			}
			if newInterval[1] > intervals[start][1] {
				intervals[start][1] = newInterval[1]
			}
		}
	} else {
		if newInterval[0] < intervals[start][0] {
			intervals[start][0] = newInterval[0]
		}

		if newInterval[1] > intervals[end][1] {
			intervals[start][1] = newInterval[1]
		} else {
			intervals[start][1] = intervals[end][1]
		}

		intervals = append(intervals[:start+1], intervals[end+1:]...)
	}

	return intervals
}
