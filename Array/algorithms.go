package Array

/*
	442. 数组中重复的数据
	https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/
*/
func findDuplicates(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		index := absInt(nums[i]) - 1
		if nums[int(index)] < 0 {
			result = append(result, nums[i])
		} else {
			nums[int(index)] *= -1
		}
	}
	return result
}

func absInt(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
	}
}
