package minimum_absolute_difference

import "slices"

func minimumAbsDifference(arr []int) [][]int {
	if len(arr) < 2 {
		return nil
	}

	slices.Sort(arr)

	pairs := [][]int{}
	lowestDiff := abs(arr[0]-arr[1]) + 1

	for i := 0; i < len(arr)-1; i++ {
		counter := dfs(arr[i+1:], arr[i])
		diff := abs(arr[i] - counter)

		if diff > lowestDiff {
			continue
		}

		var pair []int
		if arr[i] < counter {
			pair = []int{arr[i], counter}
		} else {
			pair = []int{counter, arr[i]}
		}

		if diff < lowestDiff {
			lowestDiff = diff
			pairs = [][]int{pair}
		} else {
			pairs = append(pairs, pair)
		}
	}

	return pairs
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dfs(arr []int, target int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] > target {
			end = mid - 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			return arr[mid]
		}
	}

	if start == len(arr) {
		return arr[end]
	} else {
		return arr[start]
	}
}
