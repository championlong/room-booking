package main

func maximumWealth(accounts [][]int) int {
	max, cur:= 0, 0

	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			cur += accounts[i][j]
		}
		if cur > max {
			max = cur
		}
		cur = 0
	}
	return max
}
