package main

import "math"

func beautySum(s string) int {
	ans := 0
	for i := range s {
		m := [26]int{}
		for j := i; j >= 0; j-- {
			m[s[j]-97]++
			max := 0
			min := math.MaxInt32
			for x := 0; x < 26; x++ {
				if m[x] != 0 {
					if m[x] > max {
						max = m[x]
					}
					if m[x] < min {
						min = m[x]
					}
				}
			}
			ans += max - min
		}
	}
	return ans
}
