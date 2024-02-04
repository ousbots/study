package find_median_from_data_stream

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return min
}

func (h *MinHeap) Peek() int {
	return (*h)[0]
}

type MaxHeap struct {
	MinHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

type MedianFinder struct {
	lower *MaxHeap
	upper *MinHeap
}

func Constructor() MedianFinder {
	ret := MedianFinder{lower: &MaxHeap{}, upper: &MinHeap{}}

	heap.Init(ret.lower)
	heap.Init(ret.upper)

	return ret
}

func (this *MedianFinder) AddNum(num int) {
	if this.upper.Len() > 0 && num > this.upper.Peek() {
		heap.Push(this.upper, num)
	} else {
		heap.Push(this.lower, num)
	}

	if abs(this.lower.Len()-this.upper.Len()) > 1 {
		if this.lower.Len() > this.upper.Len() {
			heap.Push(this.upper, heap.Pop(this.lower))
		} else {
			heap.Push(this.lower, heap.Pop(this.upper))
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (this *MedianFinder) FindMedian() float64 {
	if this.upper.Len() == this.lower.Len() {
		return float64(this.upper.Peek()+this.lower.Peek()) / 2.
	}

	if this.upper.Len() > this.lower.Len() {
		return float64(this.upper.Peek())
	} else {
		return float64(this.lower.Peek())
	}
}
