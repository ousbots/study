package roman_to_integer

// Convert a roman numeral string to an integer.
func romanToInt(s string) int {
	value := 0
	pos := 0

	for pos < len(s) {
		switch s[pos] {
		case 'I':
			if pos+1 < len(s) && s[pos+1] == 'V' {
				value += 4
				pos += 2
				break
			}

			if pos+1 < len(s) && s[pos+1] == 'X' {
				value += 9
				pos += 2
				break
			}

			value += 1
			pos += 1
			break

		case 'V':
			value += 5
			pos += 1
			break

		case 'X':
			if pos+1 < len(s) && s[pos+1] == 'L' {
				value += 40
				pos += 2
				break
			}

			if pos+1 < len(s) && s[pos+1] == 'C' {
				value += 90
				pos += 2
				break
			}

			value += 10
			pos += 1
			break

		case 'L':
			value += 50
			pos += 1
			break

		case 'C':
			if pos+1 < len(s) && s[pos+1] == 'D' {
				value += 400
				pos += 2
				break
			}

			if pos+1 < len(s) && s[pos+1] == 'M' {
				value += 900
				pos += 2
				break
			}

			value += 100
			pos += 1
			break

		case 'D':
			value += 500
			pos += 1
			break

		case 'M':
			value += 1000
			pos += 1
			break
		}
	}

	return value
}
