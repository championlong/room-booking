package main

func getLastMoment(n int, left []int, right []int) int {
	if n == 0 || (len(left) == 0 && len(right) == 0) {
		return 0
	}
	res := 0
	for i := 0; i < len(left); i++ {
		if left[i] > res {
			res = left[i]
		}
	}
	for i := 0; i < len(right); i++ {
		if n - right[i] > res {
			res = n - right[i]
		}
	}

	return res
}
