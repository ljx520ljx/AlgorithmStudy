package DoublePointer

//27. 移除元素
//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。
/*
算法思路:
使用双指针,一个指针指向当前不重复的元素,另一个指针遍历数组,当遇到不重复的元素时,将其赋值给第一个指针指向的位置,然后第一个指针向后移动一位。
*/
func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v != val {
			nums[index] = v
			index++
		}
	}
	return index
}
