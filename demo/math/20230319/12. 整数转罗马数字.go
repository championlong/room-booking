package main

func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	for i := 0; i < len(value); i++ {
		for num >= value[i] {
			num -= value[i]
			res += roman[i]
		}
	}
	return res
}
