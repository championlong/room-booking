package main

func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r >= 0 && r < row && c >= 0 && c < col && grid[r][c] == '1' {
			grid[r][c] = '0'
			dfs(r-1, c)
			dfs(r+1, c)
			dfs(r, c-1)
			dfs(r, c+1)
		}
	}
	var res int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
