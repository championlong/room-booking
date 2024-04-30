package main

import "sort"

var (
	res  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates)
	dfs(candidates, 0, target)
	return res
}

func dfs(candidates []int, start int, target int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, i, target-candidates[i])
		path = path[:len(path)-1]
	}
}
