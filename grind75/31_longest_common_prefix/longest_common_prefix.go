package longest_common_prefix

// Find the longest common prefix of a collection of words.
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	shortest := len(strs[0])
	for _, word := range strs {
		if len(word) < shortest {
			shortest = len(word)
		}
	}

	for i := 0; i < shortest; i++ {
		for j := range strs {
			if strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:shortest]
}
