package main

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	maxLen := 1
	var startIdx int
	for curLen := 2; curLen <= len(s); curLen++ {
		for left := 0; left < len(s)-curLen+1; left++ {
			right := left + curLen - 1
			if right >= len(s) || left+1 >= len(s) {
				break
			}

			if s[left] != s[right] {
				dp[left][right] = false
				continue
			}

			if left+1 <= right-1 {
				dp[left][right] = dp[left+1][right-1]
			} else {
				dp[left][right] = true
			}

			if dp[left][right] && right-left+1 > maxLen {
				maxLen = right - left + 1
				startIdx = left
			}
		}
	}

	return s[startIdx : startIdx+maxLen]
}
