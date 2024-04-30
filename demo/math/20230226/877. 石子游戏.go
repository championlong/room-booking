package main

func stoneGame(piles []int) bool {
	length := len(piles)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = piles[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[j] = max(piles[i]-dp[j], piles[j]-dp[j-1])
		}
	}
	return dp[length-1] > 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
