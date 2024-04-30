package main

func main() {
	isInterleave("aabcc", "dbbca", "aadbbbaccc")
}

//状态转移方程：f(i,j) = f(i-1,j) || f(i,j-1)
func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len+s2Len != s3Len {
		return false
	}
	f := make([][]bool, s1Len+1)
	for i := 0; i <= s1Len; i++ {
		f[i] = make([]bool, s2Len+1)
	}
	f[0][0] = true
	for i := 0; i <= s1Len; i++ {
		for j := 0; j <= s2Len; j++ {
			if i > 0 && s1[i-1] == s3[i+j-1] {
				f[i][j] = f[i-1][j]
			}
			if !f[i][j] {
				if j > 0 && s2[j-1] == s3[i+j-1] {
					f[i][j] = f[i][j-1]
				}
			}
		}
	}
	return f[s1Len][s2Len]
}
