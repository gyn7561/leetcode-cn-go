package _38

import "fmt"

//https://leetcode-cn.com/problems/count-and-say/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	result := ""
	lastChar := str[:1]
	count := 1
	for i := 1; i < len(str); i++ {
		char := str[i : i+1]
		if char == lastChar {
			count++
		} else {
			result += fmt.Sprintf("%d%s", count, lastChar)
			lastChar = char
			count = 1
		}
	}
	result += fmt.Sprintf("%d%s", count, lastChar)
	return result
}
