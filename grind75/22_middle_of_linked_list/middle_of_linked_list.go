package middle_of_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Find the middle node in the linked list.
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		fast = fast.Next
	}

	return slow
}
