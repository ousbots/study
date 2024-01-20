package climbing_stairs

var answers = make([]int, 45, 45)

// Count the number of ways to climb stairs taking one or two steps.
// Notes:
//   - DP problem, store calculated results
//   - Offset Fibonnaci sequence
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if answers[n-1] != 0 {
		return answers[n-1]
	}

	total := climbStairs(n-1) + climbStairs(n-2)
	answers[n-1] = total

	return total
}
