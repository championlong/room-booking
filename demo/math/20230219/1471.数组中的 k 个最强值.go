package main

import (
	"sort"
)

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	ans := make([]int, 0)

	mid := arr[(len(arr)-1)/2]
	l, r := 0, len(arr)-1
	for len(ans) < k {
		if mid-arr[l] > arr[r]-mid {
			ans = append(ans, arr[l])
			l++
		} else if mid-arr[l] < arr[r]-mid {
			ans = append(ans, arr[r])
			r--
		} else {
			ans = append(ans, arr[r])
			r--
		}
	}
	return ans
}
