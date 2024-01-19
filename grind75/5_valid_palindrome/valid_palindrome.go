package valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var word strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			word.WriteRune(unicode.ToLower(char))
		}
	}

	s = word.String()
	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] == ' ' {
			i++
			continue
		}

		if s[j] == ' ' {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
