package remove_duplicate_letters

func removeDuplicateLetters(s string) string {
	count := map[rune]int{}
	for _, letter := range s {
		count[letter]++
	}

	using := map[rune]bool{}
	remove := func(last, current rune) bool {
		if using[current] {
			return false
		}

		if count[last] == 0 {
			return false
		}

		return current < last
	}

	seen := make([]rune, 0, len(s))
	for _, letter := range s {
		count[letter]--

		for len(seen) > 0 && remove(seen[len(seen)-1], letter) {
			using[seen[len(seen)-1]] = false
			seen = seen[:len(seen)-1]
		}

		if !using[letter] {
			seen = append(seen, letter)
			using[letter] = true
		}
	}

	return string(seen)
}
