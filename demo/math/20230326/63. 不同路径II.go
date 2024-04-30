package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			} else if i >= 1 && j >= 1 {
				f[i][j] = f[i-1][j] + f[i][j-1]
			} else if i >= 1 {
				f[i][0] = f[i-1][0]
			} else if j >= 1 {
				f[0][j] = f[0][j-1]
			}
		}
	}

	return f[m-1][n-1]
}
