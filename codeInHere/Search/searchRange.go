package Search

import "sort"

//在排序数组中查找元素第一个位置和最后一个位置
//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//如果数组中不存在目标值 target，返回 [-1, -1],设计并实现时间复杂度为 O(log n)
/*
算法思路:二分查找
用二分查找来找这个数的左右边界
*/
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	n := len(nums)

	if n == 0 || target < nums[0] || target > nums[n-1] {
		return res
	}

	// Find left boundary
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return res
	}
	res[0] = left

	// Find right boundary
	right = n - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	res[1] = right

	return res
}

func searchRange1(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
