package _509

//https://leetcode-cn.com/problems/fibonacci-number/

func fib(N int) int {
	if N <= 1 {
		return N
	} else {
		return N + fib(N-1)
	}
}
