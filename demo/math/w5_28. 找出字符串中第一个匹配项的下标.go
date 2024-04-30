package main

import (
	"fmt"
)

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Println(strStr(haystack, needle))
}

//func strStr(haystack string, needle string) int {
//	return strings.Index(haystack, needle)
//}

// KMP
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	nextArray := getNextArray(needle)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = nextArray[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNextArray(needle string) []int {
	nextArray := make([]int, len(needle))
	j := -1
	nextArray[0] = j

	for i := 1; i < len(needle); i++ {
		for j >= 0 && needle[i] != needle[j+1] {
			j = nextArray[j]
		}
		if needle[i] == needle[j+1] {
			j++
		}
		nextArray[i] = j
	}
	return nextArray
}
