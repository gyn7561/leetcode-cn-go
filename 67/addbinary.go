package _67

import "strconv"

//https://leetcode-cn.com/problems/add-binary/

func parseString(num string) []int {
	nums := make([]int, len(num))
	for i := range num {
		if num[i:i+1] == "1" {
			nums[len(nums)-1-i] = 1
		} else {
			nums[len(nums)-1-i] = 0
		}
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
		result[i] = number % 2
		add = number / 2
	}
	if add > 0 {
		result = append(result, add)
	}
	return result
}

func addBinary(a string, b string) string {

	result := arrayAdd(parseString(a), parseString(b))
	strResult := ""
	for i := range result {
		strResult += strconv.Itoa(result[len(result)-1-i])
	}
	return strResult
}
