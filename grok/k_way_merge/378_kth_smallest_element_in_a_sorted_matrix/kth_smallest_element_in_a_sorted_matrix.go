package kth_smallest_element_in_a_sorted_matrix

import "container/heap"

type MaxHeap []int

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Less(a, b int) bool { return m[a] > m[b] }
func (m MaxHeap) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }
func (m *MaxHeap) Push(x any)        { *m = append(*m, x.(int)) }

func (m *MaxHeap) Pop() any {
	max := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]

	return max
}

func kthSmallest(matrix [][]int, k int) int {
	max := &MaxHeap{}
	heap.Init(max)

	for _, row := range matrix {
		for _, value := range row {
			heap.Push(max, value)

			if max.Len() > k {
				heap.Pop(max)
			}
		}
	}

	return heap.Pop(max).(int)
}
