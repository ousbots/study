package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current, last := head.Next, head
	for current != nil {
		temp := current.Next
		current.Next = last
		last = current
		current = temp
	}

	head.Next = nil

	return last
}
