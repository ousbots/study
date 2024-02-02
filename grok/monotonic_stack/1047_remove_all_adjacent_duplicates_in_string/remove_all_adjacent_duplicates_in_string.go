package remove_all_adjacent_duplicates_in_string

func removeDuplicates(s string) string {
	letters := []rune{}
	for _, letter := range s {
		if len(letters) > 0 && letter == letters[len(letters)-1] {
			letters = letters[:len(letters)-1]
		} else {
			letters = append(letters, letter)
		}
	}

	return string(letters)
}
