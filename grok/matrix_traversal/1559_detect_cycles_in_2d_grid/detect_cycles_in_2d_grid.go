package detect_cycles_in_2d_grid

func containsCycle(grid [][]byte) bool {
	seen := make([][]bool, len(grid))
	for i := range grid {
		seen[i] = make([]bool, len(grid[i]))
	}

	for i := range grid {
		for j := range grid[i] {
			if seen[i][j] == true {
				continue
			}

			if check(grid, seen, i, j, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func check(grid [][]byte, seen [][]bool, i, j, lasti, lastj, count int) bool {
	if grid[i][j] != grid[lasti][lastj] {
		return false
	}

	if count > 4 && seen[i][j] {
		return true
	}

	seen[i][j] = true

	if i > 0 && i-1 != lasti {
		if check(grid, seen, i-1, j, i, j, count+1) {
			return true
		}
	}

	if i < len(grid)-1 && i+1 != lasti {
		if check(grid, seen, i+1, j, i, j, count+1) {
			return true
		}
	}

	if j > 0 && j-1 != lastj {
		if check(grid, seen, i, j-1, i, j, count+1) {
			return true
		}
	}

	if j < len(grid[i])-1 && j+1 != lastj {
		if check(grid, seen, i, j+1, i, j, count+1) {
			return true
		}
	}

	return false
}
