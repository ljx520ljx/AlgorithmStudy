package SortAlgorithm

//最短无序连续子数组
//给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//请你找出符合题意的 最短 子数组，并输出它的长度。
/*
算法思路:拓展左右边界
初始化 left 和 right 为 -1
找到最左侧不符合升序的位置 (left)
找到最右侧不符合升序的位置 (right)
处理数组已是升序的情况
计算边界里的最小值和最大值
此时左边界左边和右边界右边都是符合升序的
根据区间内的最大最小值扩展 left 和 right 区间
时间复杂度O(n)
*/

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := -1, -1

	// 找到左侧第一个不符合升序的位置
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			left = i
			break
		}
	}

	// 找到右侧第一个不符合升序的位置
	for i := n - 1; i > 0; i-- {
		if nums[i] < nums[i-1] {
			right = i
			break
		}
	}

	// 如果整个数组已经是升序的，返回 0
	if left == -1 || right == -1 {
		return 0
	}

	// 在[left, right]区间内，找到最小值和最大值
	minVal, maxVal := nums[left], nums[right]
	for i := left; i <= right; i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	// 扩展[left, right]的区间，直到包括所有大于minVal和小于maxVal的元素
	for i := 0; i < left; i++ {
		if nums[i] > minVal {
			left = i
			break
		}
	}

	for i := n - 1; i > right; i-- {
		if nums[i] < maxVal {
			right = i
			break
		}
	}

	return right - left + 1
}
