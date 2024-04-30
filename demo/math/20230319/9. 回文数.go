package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; {
		for j := len(str) - 1; j > len(str)/2-1; j-- {
			if str[i] != str[j] {
				return false
			}
			i++
		}
	}
	return true
}
