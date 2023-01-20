func dfs(i int, j int, grid *[][]byte) {
	if (i < 0 || j < 0 || i >= len((*grid)) || j >= len((*grid)[0]) || (*grid)[i][j] == '0') {
		return
	}
	
	(*grid)[i][j] = '0'
	dfs(i, j - 1, grid)
	dfs(i, j + 1, grid)
	dfs(i - 1, j, grid)
	dfs(i + 1, j, grid)
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len (grid[0])
	islands := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (string(grid[i][j]) == "1") {
				dfs(i, j, &grid)
				islands++
			}
		}
	}
	return islands
}

