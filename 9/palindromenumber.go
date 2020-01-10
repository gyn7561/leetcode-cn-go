package _9

//https://leetcode-cn.com/problems/palindrome-number/
//回文数

func reverse(x int) int {
	num := x
	result := 0
	for {
		result = result*10 + num%10
		num = num / 10
		if num == 0 {
			break
		}
	}
	//溢出特殊处理
	if result > 2147483647 || result < -2147483647 {
		return 0
	}
	return result
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//复用翻转函数，翻转后对比一致即为回文
	return reverse(x) == x
}
