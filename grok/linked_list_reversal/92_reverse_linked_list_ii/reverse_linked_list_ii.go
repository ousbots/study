package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	var link, start, stop *ListNode

	// find the left and right positions.
	current := head
	count := 1
	for count <= right {
		if count+1 == left {
			link = current
		}

		if count == left {
			start = current
		}

		if count == right {
			stop = current
		}

		current = current.Next
		count++
	}

	// reverse the sublist.
	current = start.Next
	last := start
	for last != stop {
		temp := current.Next
		current.Next = last
		last = current
		current = temp
	}

	// re-link the reversed sublist.
	start.Next = current
	if link != nil {
		link.Next = stop
	} else {
		return stop
	}

	return head
}
