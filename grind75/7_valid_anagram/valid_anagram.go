package valid_anagram

func isAnagram(s, t string) bool {
	letters := map[rune]int{}

	for _, letter := range s {
		letters[letter]++
	}

	for _, letter := range t {
		letters[letter]--
	}

	for _, val := range letters {
		if val != 0 {
			return false
		}
	}

	return true
}
