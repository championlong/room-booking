package main

import "math"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1)
	sum := 0

	for i := n - 1; i >= 0; i-- {
		dp[i] = math.MinInt32
		sum += stoneValue[i]
		for j := i; j < i+3 && j < n; j++ {
			dp[i] = max(dp[i], sum-dp[j+1])
		}
	}

	alice := dp[0]
	bob := sum - dp[0]
	if alice > bob {
		return "Alice"
	} else if alice < bob {
		return "Bob"
	} else {
		return "Tie"
	}
}
