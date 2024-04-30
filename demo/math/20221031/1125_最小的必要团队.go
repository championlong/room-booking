package main

import "math"

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	dp := make([]int, 1<<len(req_skills))
	str2id := map[string]int{}
	for i, skill := range req_skills {
		str2id[skill] = i
	}

	for i := range dp {
		dp[i] = math.MaxInt32
	}

	peopleNum := make([]int, len(people))
	parent := map[int][]int{}
	for i, p := range people {
		sums := 0
		for _, s := range p {
			sums += 1 << str2id[s]
		}
		peopleNum[i] = sums
		dp[sums] = 1
		parent[sums] = []int{-1, i}
	}

	for i := range dp {
		if dp[i] != math.MaxInt32 {
			for j, sums := range peopleNum {
				if dp[sums|i] > dp[i]+1 {
					dp[sums|i] = dp[i] + 1
					parent[sums|i] = []int{i, j}
				}
			}
		}
	}

	ret := make([]int, 0)
	for start := len(dp) - 1; ; start = parent[start][0] {
		ret = append(ret, parent[start][1])
		if parent[start][0] == -1{
			break
		}
	}
	return ret
}
