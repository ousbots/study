package maximum_length_of_pair_chain

import "sort"

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}

	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a][1] < pairs[b][1]
	})

	count, last := 0, -99999

	for _, pair := range pairs {
		if pair[0] > last {
			count++
			last = pair[1]
		}
	}

	return count
}
