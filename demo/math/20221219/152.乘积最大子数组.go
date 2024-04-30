package _0221219

import "math"

func maxProduct(nums []int) int {
	n := len(nums)
	ans := math.MinInt
	for i := 0; i < n; i++ {
		tmp := nums[i]
		ans = max(ans, tmp)
		for j := i + 1; j < n; j++ {
			tmp *= nums[j]
			ans = max(ans, tmp)
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
