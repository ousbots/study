package ipo

import (
	"container/heap"
	"sort"
)

type Project struct {
	capital int
	profit  int
}

type ProfitQueue []int

func (p ProfitQueue) Len() int           { return len(p) }
func (p ProfitQueue) Less(i, j int) bool { return p[i] > p[j] }
func (p ProfitQueue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *ProfitQueue) Push(x any) {
	*p = append(*p, x.(int))
}

func (p *ProfitQueue) Pop() any {
	max := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]

	return max
}

func findMaximizedCapital(k, w int, profits, capital []int) int {
	if len(profits) == 0 || len(capital) == 0 || len(profits) != len(capital) {
		return -1
	}

	projects := []Project{}
	for i := range profits {
		projects = append(projects, Project{capital: capital[i], profit: profits[i]})
	}

	sort.Slice(projects, func(i, j int) bool { return projects[i].capital < projects[j].capital })

	queue := &ProfitQueue{}
	heap.Init(queue)

	total := w
	for ; k > 0; k-- {
		i := 0
		for ; i < len(projects); i++ {
			if projects[i].capital > total {
				break
			}

			heap.Push(queue, projects[i].profit)
		}
		projects = projects[i:]

		if queue.Len() == 0 {
			break
		}

		total += heap.Pop(queue).(int)
	}

	return total
}
