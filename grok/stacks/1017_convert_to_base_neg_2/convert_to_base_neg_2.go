package convert_to_base_neg_2

// Convert an integer to base -2 representation.
func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}

	output := []rune{}
	for n != 0 {
		rem := n % -2

		if rem == 0 {
			output = append([]rune{'0'}, output...)
		} else {
			output = append([]rune{'1'}, output...)
		}

		n = n / -2

		// Must use divisors that result in positive remainders. Since the remainder is either 1 or -1,
		// this means that when there is a negative remainder, the next divisor should be used, meaning n
		// should be n+1.
		if rem < 0 {
			n++
		}
	}

	return string(output)
}
