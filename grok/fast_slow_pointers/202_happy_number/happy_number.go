package happy_number

// Determine if a number is happy. A happy number is one that when repeatedly summing the squares of
// all digits results in a cycle of 1 a non-happy number results in a cycle that does not contain 1.
func isHappy(n int) bool {
	fast := n
	slow := n

	for {
		slow = calc(slow)
		fast = calc(calc(fast))

		if fast == slow {
			break
		}
	}

	return slow == 1
}

func calc(n int) int {
	digits := []int{}
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	output := 0
	for _, val := range digits {
		output += val * val
	}

	return output
}
