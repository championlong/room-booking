package main

import "strings"

func wordBreak2(s string, wordDict []string) []string {
	ans := []string{}
	dfs2(s, []string{}, &ans, &wordDict)
	return ans
}

func dfs2(str string, words []string, ans, wordDict *[]string) {
	if len(str) == 0 {
		*ans = append(*ans, strings.Join(words, " "))
		return
	}
	for _, w := range *wordDict {
		if strings.Index(str, w) == 0 {
			dfs2(str[len(w):], append(words, w), ans, wordDict)
		}
	}
}
