package main

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	var dfs func(mi int, nj int)
	dfs = func(mi, nj int) {
		if mi < 0 || nj < 0 || mi >= m || nj >= n || board[mi][nj] == 'X' || board[mi][nj] == '#' {
			return
		}
		board[mi][nj] = '#'
		dfs(mi+1, nj)
		dfs(mi, nj+1)
		dfs(mi-1, nj)
		dfs(mi, nj-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || j == 0 || j == n-1 || i == m-1
			if isEdge && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}
