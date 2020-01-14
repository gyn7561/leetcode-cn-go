package _43

import "strconv"

//https://leetcode-cn.com/problems/multiply-strings/

var numMap = make(map[uint8]int)

func initMap() {
	s := "0123456789"
	for i := range s {
		numMap[s[i]] = i
	}
}

func parseString(num string) []int {
	//个位在前方便计算
	nums := make([]int, len(num))
	for i := range num {
		nums[len(nums)-1-i] = numMap[num[i]]
	}
	return nums
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//模拟乘法
func arrayMultiply(nums []int, n int, zeroCount int) []int {
	result := make([]int, len(nums))
	add := 0 //进位
	for i := range nums {
		num := nums[i]*n + add
		result[i] = num % 10
		add = num / 10
	}
	if add > 0 {
		result = append(result, add)
	}
	for i := 0; i < zeroCount; i++ {
		result = append([]int{0}, result...)
	}
	return removeZeroPrefix(result)
}

func removeZeroPrefix(nums []int) []int {
	for {
		if len(nums) == 1 {
			break
		}

		if nums[len(nums)-1] == 0 {
			nums = nums[:len(nums)-1]
		} else {
			break
		}
	}
	return nums
}

//补0 方便计算
func addZeroPrefix(nums []int, length int) []int {
	needToAdd := length - len(nums)
	zeroArray := make([]int, needToAdd)
	return append(nums, zeroArray...)
}

func arrayAdd(a []int, b []int) []int {
	n1 := addZeroPrefix(a, max(len(a), len(b)))
	n2 := addZeroPrefix(b, max(len(a), len(b)))
	result := make([]int, len(n2))
	add := 0 //进位
	for i := 0; i < len(n1); i++ {
		number := n1[i] + n2[i] + add
		result[i] = number % 10
		add = number / 10
	}
	if add > 0 {
		result = append(result, add)
	}
	return removeZeroPrefix(result)
}

func multiply(num1 string, num2 string) string {
	initMap()
	n1 := parseString(num1)
	n2 := parseString(num2)
	sum := []int{0}
	for i := range n1 {
		sum = arrayAdd(arrayMultiply(n2, n1[i], i), sum)
	}

	strResult := ""
	for i := range sum {
		strResult += strconv.Itoa(sum[len(sum)-1-i])
	}
	return strResult

}
