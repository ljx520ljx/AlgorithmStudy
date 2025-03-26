package Search

//寻找峰值
/*
峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。
数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
且nums[i] != nums[i + 1]
*/
/*
算法思路:二分查找
二分查找：
使用二分查找的思想来寻找峰值元素。二分查找可以将搜索空间缩小一半,从而达到O(logn) 的时间复杂度。
寻找峰值的性质：
如果一个元素大于右侧相邻元素，那么在左半部分一定存在一个峰值。
如果一个元素小于右侧相邻元素，那么右半部分一定存在一个峰值。
根据以上性质，我们可以根据当前元素与其右侧元素的大小关系来决定向左还是向右继续搜索。
实现过程：
初始化两个指针 left 和 right，分别指向数组的起始和终止位置。
在每次迭代中，找到中间位置 mid，比较 nums[mid] 和 nums[mid + 1]：如果 nums[mid] > nums[mid + 1]，则说明峰值可能在左半部分，调整 right = mid。
否则，峰值在右半部分，调整 left = mid + 1。
最终，当 left == right 时，left 就指向了一个峰值的位置。
*/
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[mid+1] {
			// 必定存在峰值在左半部分
			right = mid
		} else {
			// 必定存在峰值在右半部分
			left = mid + 1
		}
	}

	return left
}
