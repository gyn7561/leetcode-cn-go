package _7

//https://leetcode-cn.com/problems/reverse-integer/
//整数反转
func reverse(x int) int {

	//每位取余累加
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
