package main

func wordBreak(s string, wordDict []string) bool {
	hash := make(map[string]bool)
	for _, v := range wordDict {
		hash[v] = true
	}
	used := make(map[int]bool)
	return dfs(s, 0, hash, used)
}

func dfs(s string, index int, hash map[string]bool, used map[int]bool) bool {
	if index == len(s) {
		return true
	}

	if val, ok := used[index]; ok {
		return val
	}

	for i := index + 1; i <= len(s); i++ {
		str := s[index:i]
		if hash[str] && dfs(s, i, hash, used) {
			used[i] = true
			return true
		}
	}
	used[index] = false
	return false
}
