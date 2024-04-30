package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	degrees := make([]int, n)
	m := make(map[int][]int)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
		degrees[edge[0]]++
		degrees[edge[1]]++
	}

	var queue []int
	for i, degree := range degrees {
		if degree == 1 {
			queue = append(queue, i)
		}
	}

	var ans []int
	for len(queue) > 0 {
		var level []int
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			level = append(level, cur)
			for _, nei := range m[cur] {
				degrees[nei]--
				if degrees[nei] == 1 {
					queue = append(queue, nei)
				}
			}
		}
		queue = queue[size:]
		ans = level
	}
	return ans
}
