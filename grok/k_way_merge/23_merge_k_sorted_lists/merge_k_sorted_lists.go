package merge_k_sorted_lists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinNode []*ListNode

func (m MinNode) Len() int           { return len(m) }
func (m MinNode) Less(a, b int) bool { return m[a].Val < m[b].Val }
func (m MinNode) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }

func (m *MinNode) Push(x any) {
	*m = append(*m, x.(*ListNode))
}

func (m *MinNode) Pop() any {
	min := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]

	return min
}

func mergeKLists(lists []*ListNode) *ListNode {
	nodes := &MinNode{}
	heap.Init(nodes)

	for _, root := range lists {
		if root != nil {
			heap.Push(nodes, root)
		}
	}

	root := ListNode{}
	current := &root

	for nodes.Len() > 0 {
		node := heap.Pop(nodes).(*ListNode)
		current.Next = node
		current = node

		if node.Next != nil {
			heap.Push(nodes, node.Next)
		}
	}

	return root.Next
}
