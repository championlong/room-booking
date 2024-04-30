package main

import "math"

func maxProfit(prices []int) int {
	maxInt, result := math.MaxInt, 0
	num := len(prices) - 1
	for i := 0; i <= num; i++ {
		maxInt = min(maxInt, prices[i])
		result = max(prices[i]-maxInt, result)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
