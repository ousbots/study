package simplify_path

func simplifyPath(path string) string {
	dirs := []string{}

	for start := 0; start < len(path); {
		if path[start] == '/' {
			start++
			continue
		}

		for end := start; end < len(path); end++ {
			if path[end] == '/' {
				dirs = append(dirs, path[start:end])
				start = end
				break
			}

			if end == len(path)-1 {
				dirs = append(dirs, path[start:end+1])
				start = end + 1
				break
			}
		}
	}

	filtered := []string{}
	for i := range dirs {
		if dirs[i] == "." {
			continue
		}

		if dirs[i] == ".." {
			if len(filtered) > 0 {
				filtered = filtered[:len(filtered)-1]
			}
			continue
		}

		filtered = append(filtered, dirs[i])
	}

	output := "/"
	for len(filtered) > 0 {
		dir := filtered[0]
		filtered = filtered[1:]
		output += dir

		if len(filtered) > 0 {
			output += "/"
		}
	}

	return output
}
