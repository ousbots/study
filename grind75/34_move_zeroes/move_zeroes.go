package move_zeroes

// Move all the zeroes in the array to the end in-place.
// Notes:
//
//	This uses two pointers to build the non-zero part of the array by assigning all non-zero
//	values to the left pointer and not incrementing the left pointer when a zero is encountered.
//	Then the left pointer to the end of the array can be assigned zeroes to "clear" the remainder
//	of the array.
func moveZeroes(nums []int) {
	left := 0

	for right := range nums {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}

	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}

// Move all the zeroes in the array to the end in-place.
// Note: bubble all zeroes to the end.
func moveZeroesSlow(nums []int) {
	//for i := range nums {
	count := 0
	for i := 0; i < len(nums)-1-count; i++ {
		if nums[i] == 0 {
			for j := i; j < len(nums)-1; j++ {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}

			count++
			i--
		}
	}
}
