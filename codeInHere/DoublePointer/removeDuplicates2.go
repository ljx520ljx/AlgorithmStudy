package DoublePointer

//80. 删除有序数组中的重复项 II
//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。
/*
算法思路:
	双指针
		index 为慢指针，记录当前不重复的元素位置
		i 为快指针，遍历数组
		当nums[i] != nums[index-1] || nums[i] != nums[index-2]时，说明当前元素不重复，将nums[i]赋值给nums[index]，index++
		当nums[i] == nums[index-1] || nums[i] == nums[index-2]时，说明当前元素重复，跳过，i++
		最后返回index
*/
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	index := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[index-1] || nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
