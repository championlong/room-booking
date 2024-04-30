package main

func main() {
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word := "ABCCED"
	exist(board,word)
}


func exist(board [][]byte, word string) bool {
	found := false
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		// 超出索引范围
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		// 走过，不能再走
		if board[i][j] == '1' {
			return
		}
		// 元素不相等
		if board[i][j] != word[k] {
			return
		}
		// 元素相等 && 长度相等，标记找到
		if k == len(word)-1 {
			found = true
			return
		}
		// 标记走过
		tmp := board[i][j]
		board[i][j] = '1'

		// 继续往后走
		dfs(i-1, j, k+1)
		dfs(i+1, j, k+1)
		dfs(i, j-1, k+1)
		dfs(i, j+1, k+1)

		// 走完之后回溯状态，以便于下一个点的统计
		board[i][j] = tmp
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0 // index of the word
			dfs(i, j, k)
		}
	}
	return found
}
