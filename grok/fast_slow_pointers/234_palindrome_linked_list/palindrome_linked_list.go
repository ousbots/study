package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(root *ListNode) bool {
	if root == nil {
		return false
	}

	seen := []int{}
	mid := root
	fast := root

	for fast != nil {
		seen = append(seen, mid.Val)

		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next

		mid = mid.Next
	}

	for mid != nil && len(seen) > 0 {
		check := seen[len(seen)-1]
		seen = seen[:len(seen)-1]

		if check != mid.Val {
			return false
		}

		mid = mid.Next
	}

	if len(seen) > 0 || mid != nil {
		return false
	}

	return true
}
