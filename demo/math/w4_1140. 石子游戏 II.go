package main

func stoneGameII(piles []int) int {
	dp := make([][]int, len(piles))
	length := len(piles)
	total := 0
	for i := length - 1; i >= 0; i-- {
		dp[i] = make([]int, len(piles)+1)
		total += piles[i]
		for M := 1; M <= length; M++ {
			if length-i <= M*2 {
				dp[i][M] = total
			} else {
				for X := 1; X <= 2*M; X++ {
					dp[i][M] = max(dp[i][M], total-dp[i+X][max(M, X)])
				}
			}
		}
	}
	return dp[0][1]
}
