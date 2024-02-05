package top_k_frequent_elements

import "container/heap"

type Count struct {
	Value int
	Count int
}

type MaxCount []Count

func (m MaxCount) Len() int           { return len(m) }
func (m MaxCount) Less(a, b int) bool { return m[a].Count > m[b].Count }
func (m MaxCount) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }
func (m *MaxCount) Push(x any)        { *m = append(*m, x.(Count)) }

func (m *MaxCount) Pop() any {
	max := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]

	return max
}

func topKFrequent(nums []int, k int) []int {
	counts := map[int]int{}

	for _, val := range nums {
		counts[val]++
	}

	top := &MaxCount{}
	heap.Init(top)

	for val, count := range counts {
		heap.Push(top, Count{Value: val, Count: count})
	}

	output := []int{}
	for ; k > 0; k-- {
		output = append(output, heap.Pop(top).(Count).Value)
	}

	return output
}
