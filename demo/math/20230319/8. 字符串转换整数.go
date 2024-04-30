package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	ans, pos := 0, 1

	for i := range s {
		if i == 0 && s[i] == '+' {
			continue
		}
		if i == 0 && s[i] == '-' {
			pos = -1
			continue
		}
		if i == 0 && (s[i] > '9' || s[i] < '0') {
			break
		}
		if s[i] < '0' || s[i] > '9' {
			break
		}
		ans = ans*10 + int(s[i]-'0')
		if ans*pos > math.MaxInt32 {
			return math.MaxInt32
		}
		if ans*pos < math.MinInt32 {
			return math.MinInt32
		}
	}
	return ans * pos
}
