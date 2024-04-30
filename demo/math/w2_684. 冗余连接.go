package main

func findRedundantConnection(edges [][]int) []int {
	parents := make([]int, len(edges)+1)
	sizes := make([]int, len(edges)+1)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		if parents[u] == 0 {
			parents[u] = u
		}
		if parents[v] == 0 {
			parents[v] = v
		}

		pu := find(u, parents)
		pv := find(v, parents)

		if pu == pv {
			return edge
		}

		if pv > pu {
			pu, pv = pv, pu
		}

		parents[pv] = pu
		sizes[pu] += sizes[pv]

	}

	return nil
}

func find(v int, parents []int) int {
	for parents[v] != v {
		parents[v] = parents[parents[v]]
		v = parents[v]
	}
	return v
}
