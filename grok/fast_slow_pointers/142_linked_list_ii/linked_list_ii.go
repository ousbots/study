package linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(root *ListNode) *ListNode {
	slow := root
	fast := root

	for {
		if fast == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next

		if fast == nil {
			return nil
		}
		fast = fast.Next

		if fast == slow {
			break
		}
	}

	start := root
	for start != slow {
		start = start.Next
		slow = slow.Next
	}

	return start
}

func detectCycleMem(root *ListNode) *ListNode {
	seen := make(map[*ListNode]struct{})

	i := 0
	for root != nil {
		if _, exists := seen[root]; exists {
			return root
		}

		seen[root] = struct{}{}
		root = root.Next
		i++
	}

	return nil
}
