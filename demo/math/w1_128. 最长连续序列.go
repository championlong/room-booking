package main

func longestConsecutive(nums []int) int {
	num := map[int]bool{}
	ans := 0
	for _, v := range nums {
		num[v] = true
	}
	for v := range num {
		if !num[v-1] {
			tmp := 0
			for num[v] {
				tmp++
				v++
			}
			if tmp > ans {
				ans = tmp
			}
		}

	}
	return ans
}
