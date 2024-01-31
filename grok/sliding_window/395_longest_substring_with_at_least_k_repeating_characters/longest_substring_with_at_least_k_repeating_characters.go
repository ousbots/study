package longest_substring_with_at_least_k_repeating_characters

func longestSubstring(s string, k int) int {
	return check([]byte(s), 0, len(s), k)
}

func check(s []byte, start, end, k int) int {
	if end-start+1 < k {
		return 0
	}

	freq := make([]int, 26)
	for i := start; i < end; i++ {
		freq[s[i]-'a'] += 1
	}

	for i := start; i < end; i++ {
		if freq[s[i]-'a'] < k {
			left := check(s, start, i, k)
			right := check(s, i+1, end, k)

			if left > right {
				return left
			} else {
				return right
			}
		}
	}

	return end - start
}

func longestSubstringBrute(s string, k int) int {
	maximum := 0

	for start := 0; start <= len(s)-k; start++ {
		seen := make([]int, 26)

		for end := start; end < len(s); end++ {
			length := end - start + 1
			seen[s[end]-'a'] += 1

			if length > maximum && seen[s[end]-'a'] >= k && valid(seen, k) {
				maximum = length
			}
		}
	}

	return maximum
}

func longestSubstringSliding(s string, k int) int {
	maximum := 0

	for length := k; length <= len(s); length++ {
		seen := make([]int, 26)

		start := 0
		end := 0

		for end < len(s) {
			if end < length {
				seen[s[end]-'a'] += 1
				end++
			} else {
				seen[s[start]-'a'] -= 1
				seen[s[end]-'a'] += 1

				start++
				end++
			}

			if end >= length && seen[s[end-1]-'a'] >= k {
				if ok := valid(seen, k); ok && length > maximum {
					maximum = length
					break
				}
			}
		}
	}

	return maximum
}

func valid(seen []int, k int) bool {
	for _, count := range seen {
		if count < k && count > 0 {
			return false
		}
	}

	return true
}
