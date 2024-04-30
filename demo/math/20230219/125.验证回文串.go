package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; {
		if !validChar(s[i]) {
			i++
			continue
		}
		if !validChar(s[j]) {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func validChar(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}
