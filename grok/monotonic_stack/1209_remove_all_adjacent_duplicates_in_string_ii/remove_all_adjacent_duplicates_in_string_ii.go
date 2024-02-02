package remove_all_adjacent_duplicates_in_string_ii

func removeDuplicates(s string, k int) string {
	letters := []rune{}

	for _, letter := range s {
		count := 1
		for i := len(letters) - 1; i >= 0 && letters[i] == letter; i-- {
			count++

			if count == k {
				break
			}
		}

		if count == k {
			letters = letters[:len(letters)-k+1]
			continue
		}

		letters = append(letters, letter)
	}

	return string(letters)
}
