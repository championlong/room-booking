package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	numStack := make([]int, 0)
	strStack := make([]string, 0)
	num := 0
	res := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n, _ := strconv.Atoi(string(c))
			num = num*10 + n
		} else if c == '[' {
			strStack = append(strStack, res)
			res = ""
			numStack = append(numStack, num)
			num = 0
		} else if c == ']' {
			repeat := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			res = str + strings.Repeat(res, repeat)
		} else {
			res = res + string(c)
		}
	}
	return res
}
