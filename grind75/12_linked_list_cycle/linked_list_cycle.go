package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// Determine if a linked list has a cycle.
func hasCycle(head *ListNode) bool {
	first := head
	second := head

	for first != nil && second != nil {
		if first != nil {
			first = first.Next
		}

		if second != nil {
			second = second.Next
		}

		if second != nil {
			second = second.Next
		}

		if first == second && first != nil && second != nil {
			return true
		}
	}

	return false
}
