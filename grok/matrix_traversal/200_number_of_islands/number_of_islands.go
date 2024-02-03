package number_of_islands

func numIslands(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				markIsland(grid, i, j)
				count++
			}
		}
	}

	return count
}

func markIsland(grid [][]byte, i, j int) {
	q := [][]int{{i, j}}

	for len(q) > 0 {
		check := q[0]
		q = q[1:]

		if grid[check[0]][check[1]] == '0' {
			continue
		}

		grid[check[0]][check[1]] = '0'

		if check[0] > 0 {
			q = append(q, []int{check[0] - 1, check[1]})
		}

		if check[0] < len(grid)-1 {
			q = append(q, []int{check[0] + 1, check[1]})
		}

		if check[1] > 0 {
			q = append(q, []int{check[0], check[1] - 1})
		}

		if check[1] < len(grid[check[0]])-1 {
			q = append(q, []int{check[0], check[1] + 1})
		}
	}
}
