package backspace_string_compare

// Compare if two strings are equal when the '#' character is interpreted as a backspace.
func backspaceCompare(s, t string) bool {
	sp := parse(s)
	tp := parse(t)

	if len(sp) != len(tp) {
		return false
	}

	for i := range sp {
		if sp[i] != tp[i] {
			return false
		}
	}

	return true
}

func parse(input string) string {
	output := []rune{}

	for _, char := range input {
		if char == '#' {
			if len(output) > 0 {
				output = output[:len(output)-1]
			}
		} else {
			output = append(output, char)
		}
	}

	return string(output)
}
