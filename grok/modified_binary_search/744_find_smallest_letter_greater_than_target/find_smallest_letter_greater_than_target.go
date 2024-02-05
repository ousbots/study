package find_smallest_letter_greater_than_target

func nextGreatestLetter(letters []byte, target byte) byte {
	start, end := 0, len(letters)-1

	for start <= end {
		mid := (start + end) / 2

		if letters[mid] > target {
			end = mid - 1
		} else if letters[mid] < target {
			start = mid + 1
		} else {
			for i := mid + 1; i < len(letters); i++ {
				if letters[i] != letters[mid] {
					return letters[i]
				}
			}

			return letters[0]
		}
	}

	if start == len(letters) {
		return letters[0]
	}

	return letters[start]
}
