package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge two sorted lists into a single sorted list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := ListNode{}
	current := &root

	for {
		if list1 == nil && list2 == nil {
			break
		}

		if list1 == nil {
			current.Next = list2
			list2 = list2.Next
			current = current.Next
			current.Next = nil
			continue
		}

		if list2 == nil {
			current.Next = list1
			list1 = list1.Next
			current = current.Next
			current.Next = nil
			continue
		}

		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
			current = current.Next
			current.Next = nil
		} else {
			current.Next = list2
			list2 = list2.Next
			current = current.Next
			current.Next = nil
		}
	}

	return root.Next
}
