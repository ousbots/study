package ransom_note

// Determine if the ransom note can be constructed using the magazine.
func canConstruct(ransomNote, magazine string) bool {
	letters := map[rune]int{}

	for _, letter := range magazine {
		letters[letter]++
	}

	for _, letter := range ransomNote {
		letters[letter]--
	}

	for _, count := range letters {
		if count < 0 {
			return false
		}
	}

	return true
}
