package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse a linked list.
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	next := head.Next
	var last *ListNode

	for current != nil {
		next = current.Next
		current.Next = last
		last = current
		current = next
	}

	return last
}
