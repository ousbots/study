package valid_parentheses

// Determine if a string of parentheses are closed in the same order they're opened.
func isValid(s string) bool {
	input := []rune{}

	for _, char := range s {
		input = append(input, char)
	}

	open := []rune{}

	for _, symbol := range input {
		if symbol == '(' || symbol == '[' || symbol == '{' {
			open = append(open, symbol)
			continue
		}

		if len(open) == 0 {
			return false
		}

		last := open[len(open)-1]
		open = open[:len(open)-1]

		if symbol == ')' && last == '(' {
			continue
		}

		if symbol == ']' && last == '[' {
			continue
		}

		if symbol == '}' && last == '{' {
			continue
		}

		return false
	}

	if len(open) > 0 {
		return false
	}

	return true
}
