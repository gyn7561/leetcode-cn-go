package _1

//https://leetcode-cn.com/problems/two-sum/
//1. 两数之和

func twoSum(nums []int, target int) []int {
	//最简单的双循环对比
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				if nums[i]+nums[j] == target {
					result[0] = i
					result[1] = j
					return result
				}
			}
		}
	}
	return result
}
