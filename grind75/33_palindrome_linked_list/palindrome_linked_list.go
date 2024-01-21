package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Determine if a linked list is a palindrome.
func isPalindromeMem(head *ListNode) bool {
	vals := []int{}

	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	start := 0
	end := len(vals) - 1

	for start <= end {
		if vals[start] != vals[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// Determine if a linked list is a palindrome.
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	fast := head
	mid := head
	seen := []*ListNode{}
	for fast != nil {
		seen = append(seen, mid)

		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		mid = mid.Next
	}

	for len(seen) > 0 {
		if mid == nil {
			return false
		}

		check := seen[len(seen)-1]
		seen = seen[:len(seen)-1]

		if check.Val != mid.Val {
			return false
		}

		mid = mid.Next
	}

	return true
}
