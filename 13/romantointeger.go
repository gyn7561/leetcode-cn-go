package _13

//https://leetcode-cn.com/problems/roman-to-integer/submissions/
//罗马数字转整数
func romanToInt(s string) int {
	numberMap := make(map[string]int)
	numberMap["I"] = 1
	numberMap["V"] = 5
	numberMap["X"] = 10
	numberMap["L"] = 50
	numberMap["C"] = 100
	numberMap["D"] = 500
	numberMap["M"] = 1000
	var strLen = len(s)
	result := 0
	lastNumber := 0
	for v := range s {
		num := numberMap[string(s[strLen-v-1])]
		if lastNumber <= num {
			result += num
		} else {
			result -= num
		}
		lastNumber = num
	}
	println(result)
	return result
}
