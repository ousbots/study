package counting_bits

// Build an array of length n+1 where each element i is the number of ones in the binary
// representation of i.
// Note: I just looked up a popcount algorithm
func countBits(n int) []int {
	output := []int{}

	for i := 0; i <= n; i++ {
		val := i
		count := 0

		for val != 0 {
			val = val & (val - 1)
			count++
		}

		output = append(output, count)
	}

	return output
}
