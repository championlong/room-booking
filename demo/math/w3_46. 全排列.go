package main

var (
	st []bool
)

func permute(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))
	dfs2(nums, 0)
	return res
}

func dfs2(nums []int, cur int) {
	if cur == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if !st[i] {
			path = append(path, nums[i])
			st[i] = true
			dfs2(nums, cur+1)
			st[i] = false
			path = path[:len(path)-1]
		}
	}
}
