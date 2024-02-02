package longest_palindrome

func longestPalindrome(s string) int {
	letters := map[rune]int{}

	for _, letter := range s {
		letters[letter]++
	}

	total := 0
	odd := false

	for _, count := range letters {
		if count%2 == 0 {
			total += count
		} else {
			total += count - 1
			odd = true
		}
	}

	if odd {
		total++
	}

	return total
}
