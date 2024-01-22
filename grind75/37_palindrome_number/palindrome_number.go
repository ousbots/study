package palindrome_number

// Determine if a number is a palindrome.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := []int{}
	for x >= 1 {
		digits = append([]int{x % 10}, digits...)
		x /= 10
	}

	first := 0
	last := len(digits) - 1

	for first < last {
		if digits[first] != digits[last] {
			return false
		}

		first++
		last--
	}

	return true
}
