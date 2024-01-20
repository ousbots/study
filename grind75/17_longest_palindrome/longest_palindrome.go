package longest_palindrome

func isOdd(n int) bool {
	if n%2 == 1 {
		return true
	}

	return false
}

// Determine the longest palindrome that can be created using the letters of the given string.
func longestPalindrome(s string) int {
	letters := map[rune]int{}
	for _, letter := range s {
		letters[letter]++
	}

	odd := false
	total := 0
	for _, count := range letters {
		if isOdd(count) {
			odd = true
			total += count - 1
		} else {
			total += count
		}
	}

	if odd {
		total++
	}

	return total
}
