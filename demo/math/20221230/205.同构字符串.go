package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	l := len(s)
	m1 := map[uint8]int{}
	m2 := map[uint8]int{}
	for i := 0; i < l; i++ {
		c1, ok1 := m1[s[i]]
		c2, ok2 := m2[t[i]]
		if ok1 != ok2 || c1 != c2 {
			return false
		}
		if !ok1 {
			m1[s[i]] = i
			m2[t[i]] = i
		}
	}
	return true
}
