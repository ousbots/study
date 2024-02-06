package find_k_closest_elements

import (
	"container/heap"
	"slices"
)

type DiffHeap struct {
	Target int
	Q      []int
}

func (d DiffHeap) Len() int      { return len(d.Q) }
func (d DiffHeap) Swap(a, b int) { d.Q[a], d.Q[b] = d.Q[b], d.Q[a] }
func (d *DiffHeap) Push(x any)   { d.Q = append(d.Q, x.(int)) }

func (d DiffHeap) Less(a, b int) bool {
	diffA := abs(d.Q[a] - d.Target)
	diffB := abs(d.Q[b] - d.Target)

	if diffA == diffB {
		return d.Q[a] < d.Q[b]
	}

	return diffA < diffB
}

func (d *DiffHeap) Pop() any {
	min := d.Q[len(d.Q)-1]
	d.Q = d.Q[:len(d.Q)-1]

	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func findClosestElements(arr []int, k, x int) []int {
	closest := &DiffHeap{Target: x}
	heap.Init(closest)

	for _, val := range arr {
		heap.Push(closest, val)
	}

	output := make([]int, k)
	for i := 0; i < k; i++ {
		output[i] = heap.Pop(closest).(int)
	}
	slices.Sort(output)

	return output
}
