package DoublePointer

//26. 删除有序数组中的重复项
/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
*/
/*
算法思路:
使用双指针，一个指针指向当前不重复的元素，另一个指针遍历数组，当遇到不重复的元素时，将其赋值给第一个指针指向的位置，然后第一个指针向后移动一位。
*/
func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
