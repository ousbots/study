package maximum_number_of_balloons

func maxNumberOfBalloons(test string) int {
	letters := map[rune]int{}

	for _, letter := range test {
		if letter == 'a' || letter == 'b' || letter == 'l' || letter == 'o' || letter == 'n' {
			letters[letter]++
		}
	}

	letters['l'] /= 2
	letters['o'] /= 2

	min := -1
	found := 0
	for _, count := range letters {
		found++
		if min == -1 || count < min {
			min = count
		}
	}

	if found < 5 {
		return 0
	}

	return min
}
