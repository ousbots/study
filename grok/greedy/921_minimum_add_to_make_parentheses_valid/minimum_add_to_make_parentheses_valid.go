package minimum_add_to_make_parentheses_valid

func minAddToMakeValid(s string) int {
	chars := []rune(s)
	start, end, pairs := 0, 1, 0

	for start < end && end < len(chars) {
		if chars[start] == '(' && chars[end] == ')' {
			pairs++
			start++
			end++
			continue
		}

		if chars[start] != '(' {
			start++
		}

		if chars[end] != ')' {
			end++
		}

		if start == end {
			end++
		}
	}

	return len(chars) - (2 * pairs)
}
