package first_bad_version

var badVersion int

func isBadVersion(version int) bool {
	if version >= badVersion {
		return true
	}

	return false
}

// Determine the first bad version using an "api" call to isBadVersion.
func firstBadVersion(n int) int {
	start := 0
	end := n

	for start <= end {
		mid := (start + end) / 2

		if start == end {
			return start
		}

		if !isBadVersion(mid) {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return -1
}
