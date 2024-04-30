package main

func findLUSlength(strs []string) int {
	res := -1
	for i := range strs {
		s := strs[i]
		sIsSubStrT := false
		for j := range strs {
			if j == i {
				continue
			}
			t := strs[j]
			if isSubStr(s, t) {
				sIsSubStrT = true
			}
		}
		if !sIsSubStrT && res < len(s) {
			res = len(s)
		}
	}
	return res
}

func isSubStr(s, t string) bool {
	pS := 0
	for pT := range t {
		if t[pT] == s[pS] {
			pS++
			if pS == len(s) {
				return true
			}
		}
	}
	return false
}
