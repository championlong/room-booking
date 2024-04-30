package main

func minNumberOfFrogs(croakOfFrogs string) int {
	counter := map[int]int{}
	runeMap := map[rune]int{'c': 1, 'r': 2, 'o': 3, 'a': 4, 'k': 5}
	for _, r := range croakOfFrogs {
		c := runeMap[r]
		if counter[c-1] <= 0 {
			if c != 1 {
				return -1
			}
			counter[0] = 1
		}
		counter[c%5] += 1
		counter[c-1] -= 1
	}
	for i := 1; i <= 4; i++ {
		if counter[i] > 0 {
			return -1
		}
	}
	return counter[0]
}
