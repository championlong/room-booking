package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "25525511135"
	result := restoreIpAddresses(s)
	fmt.Println(result)
}

var ans []string

func restoreIpAddresses(s string) []string {
	ans = make([]string, 0)
	backtracking(s, []string{}, 0, 0)
	return ans
}

func backtracking(s string, path []string, startIdx, depth int) {
	if depth == 4 && startIdx == len(s) && len(path) == 4 {
		ip := strings.Join(path, ".")
		ans = append(ans, ip)
		return
	}
	for i:=startIdx;i<len(s) && i<startIdx+3; i++{
		if i - startIdx == 2{
			if s[startIdx:i+1] > "255"{
				continue
			}
		}
		path = append(path, s[startIdx:i+1])
		backtracking(s, path, i+1, depth+1)
		path = path[:len(path)-1]
		if s[startIdx] == '0' {
			break
		}
	}
}
