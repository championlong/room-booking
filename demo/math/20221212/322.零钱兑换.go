package main

func coinChange(coins []int, amount int) int {
	maxVal := amount + 1
	dp := make([]int, maxVal)
	for i := 0; i < maxVal; i++ {
		dp[i] = maxVal
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == maxVal {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
