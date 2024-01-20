package add_binary

// Add two binary strings.
func addBinary(a, b string) string {
	aRunes := []rune(a)
	bRunes := []rune(b)

	ai := len(aRunes) - 1
	bi := len(bRunes) - 1

	output := []rune{}

	carry := false
	for ai >= 0 || bi >= 0 {
		if ai >= 0 && bi >= 0 {
			val, remain := add(aRunes[ai], bRunes[bi], carry)
			output = append([]rune{val}, output...)
			carry = remain
		} else if ai >= 0 {
			val, remain := add(aRunes[ai], '0', carry)
			output = append([]rune{val}, output...)
			carry = remain
		} else {
			val, remain := add(bRunes[bi], '0', carry)
			output = append([]rune{val}, output...)
			carry = remain
		}

		ai--
		bi--
	}

	if carry {
		output = append([]rune{'1'}, output...)
	}

	return string(output)
}

func add(a, b rune, carry bool) (rune, bool) {
	if a == '0' {
		if b == '0' {
			if carry {
				return '1', false
			} else {
				return '0', false
			}
		} else {
			if carry {
				return '0', true
			} else {
				return '1', false
			}
		}
	} else {
		if b == '0' {
			if carry {
				return '0', true
			} else {
				return '1', false
			}
		} else {
			if carry {
				return '1', true
			} else {
				return '0', true
			}
		}
	}
}
