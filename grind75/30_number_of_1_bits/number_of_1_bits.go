package number_of_1_bits

// Determine the number of ones in the binary representation of the number.
// Note: looked up a hamming weight algorithm.
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += 1
		num &= num - 1
	}

	return count
}
