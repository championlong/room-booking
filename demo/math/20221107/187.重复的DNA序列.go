package main

import "fmt"

func main() {
	array := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	//array := findRepeatedDnaSequences("AAAAAAAAAAA")
	fmt.Println(array)
}
/*
DNA序列由一系列核苷酸组成，缩写为'A','C','G'和'T'。

例如，"ACGAATTCCG"是一个 DNA序列 。
在研究 DNA 时，识别 DNA 中的重复序列非常有用。

给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的长度为10的序列(子字符串)。你可以按 任意顺序 返回答案。
 */
func findRepeatedDnaSequences(s string) []string {
	cacheMap := make(map[string]int)
	var result []string
	for i := 0; i <= len(s)-10; i++ {
		key := s[i : i+10]
		cacheMap[key] ++
		if cacheMap[key] > 1 {
			result = append(result, key)
		}
	}
	return result
}


/*
	leetCode官方题解：
	A 表示为二进制 00；
	C 表示为二进制 01；
	G 表示为二进制 10；
	T 表示为二进制 11。

 */
const L = 10
var bin = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences1(s string) (ans []string) {
	n := len(s)
	if n <= L {
		return
	}
	x := 0
	for _, ch := range s[:L-1] {
		x = x<<2 | bin[byte(ch)]
	}
	//x: 000000000001010101
	//1<<(L*2) - 1: 11111111111111111111
	cnt := map[int]int{}
	for i := 0; i <= n-L; i++ {
		x = (x<<2 | bin[s[i+L-1]]) & (1<<(L*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i:i+L])
		}
	}
	return ans
}

