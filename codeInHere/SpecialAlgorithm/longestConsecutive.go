package SpecialAlgorithm

//最长连续序列
//一个未排序的整数数组,nums找出数字连续(如:1,2,3)的最长序列（不要求序列元素在原数组中连续）的长度。
//要求时间复杂度O(n).
/*
算法思路:hash表
枚举数组中的每个数 x，考虑以其为起点，不断尝试匹配 x+1,x+2,⋯ 是否存在，假设最长匹配到了 x+y，
那么以 x 为起点的最长连续序列即为 x,x+1,x+2,⋯,x+y，其长度为 y+1，我们不断枚举并更新答案
*/

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0
	for num := range numSet {
		// 如果这个数字是一个序列的起始点
		if !numSet[num-1] {
			currentNum := num
			currentLen := 1

			// 向后查找连续的数字
			for numSet[currentNum+1] {
				currentNum++
				currentLen++
			}
			maxLen = max(maxLen, currentLen)
		}
	}

	return maxLen
}
