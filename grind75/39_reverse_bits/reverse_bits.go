package reverse_bits

// Reverse the bits in a 32 bit number.
func reverseBits(num uint32) uint32 {
	total := uint32(0)
	for i := 0; i < 32; i++ {
		total = total << 1
		total += num & 1
		num = num >> 1
	}

	return total
}
