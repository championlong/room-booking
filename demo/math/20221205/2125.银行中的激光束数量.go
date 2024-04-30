package main

func numberOfBeams(bank []string) int {
	var sum, pre, cur int
	for _, r := range bank {
		cur = 0
		for _, b := range r {
			if b == '1' {
				cur++
			}
		}
		if cur > 0 {
			if pre > 0 {
				sum += cur * pre
			}
			pre = cur
		}
	}
	return sum
}
