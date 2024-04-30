package main

import "strings"

var res [][]string

func solveNQueens(n int) [][]string {
	res = [][]string{}
	cur := make([][]string, n)
	for i := 0; i < n; i++ {
		cur[i] = make([]string, n)
		for j := 0; j < n; j++ {
			cur[i][j] = "."
		}
	}
	dfs(0, n, 0, cur)
	return res
}

func dfs(k, n, curRow int, cur [][]string) {
	if k == n {
		add(cur)
		return
	}
	for j := 0; j < n; j++ {
		if cur[curRow][j] == "." {
			if check(curRow, j, cur) {
				cur[curRow][j] = "Q"
				dfs(k+1, n, curRow+1, cur)
				cur[curRow][j] = "."

			}
		}
	}
	return
}

func add(cur [][]string) {
	var temp []string
	for _, obj := range cur {
		temp = append(temp, strings.Join(obj, ""))
	}
	res = append(res, temp)
}

func check(row, col int, cur [][]string) bool {

	for i := 0; i < len(cur); i++ {
		if cur[row][i] == "Q" || cur[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if cur[i][j] == "Q" {
			return false
		}
		i--
		j--
	}
	for i, j := row-1, col+1; i >= 0 && j < len(cur); {
		if cur[i][j] == "Q" {
			return false
		}
		i--
		j++
	}
	return true
}
