func shortestPathBinaryMatrix(grid [][]int) int {
    n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	queue := [][]int{{0, 0}}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	grid[0][0] = 1

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		row, col := curr[0], curr[1]

		if row == n-1 && col == n-1 {
			return grid[row][col]
		}

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < n && newCol >= 0 && newCol < n && grid[newRow][newCol] == 0 {
				queue = append(queue, []int{newRow, newCol})
				grid[newRow][newCol] = grid[row][col] + 1
			}
		}
	}

	return -1
}
