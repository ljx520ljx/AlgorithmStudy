package SpecialAlgorithm

//27. 移除元素
//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。
/*
算法思路:指针遍历
遍历数组,遇到不等于val的值就把它放到index位置上,index++
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
