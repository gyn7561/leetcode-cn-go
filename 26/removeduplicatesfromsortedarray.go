package _26

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {

	index := 0
	for {

		length := len(nums)

		if length <= 1 {
			return len(nums)
		}

		end := true
		for i := index; i < length-1; i++ {
			if nums[i] == nums[i+1] {
				nums = append(nums[:i+1], nums[i+2:]...)
				end = false
				break
			} else {
				index = i
			}
		}

		if end {
			break
		}
	}
	return len(nums)
}
