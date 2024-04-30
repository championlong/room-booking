package main

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] > '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
