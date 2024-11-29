package Search

// 搜索旋转排序数组
// 例如，[0,1,2,4,5,6,7]在下标3处经旋转后可能变为[4,5,6,7,0,1,2]。
// 给你旋转后的数组nums和一个整数target，如果nums中存在这个目标值target,
//则返回它的下标，否则返回-1。
// 你必须设计一个时间复杂度为O(log n)的算法解决此问题。
/*
算法思路:二分查找
二分查找：
由于数组是部分有序的，我们可以利用二分查找的思想，将时间复杂度控制在 O(log⁡n)O(\log n)O(logn)。
每次选择中间元素 mid，并判断哪一半是有序的。
判断哪一半是有序的：
如果 nums[left] <= nums[mid]，说明左半部分是有序的。如果目标值 target 在 nums[left] 和 nums[mid] 之间，则在左半部分查找，否则在右半部分查找。
如果 nums[mid] < nums[right]，说明右半部分是有序的。如果目标值 target 在 nums[mid] 和 nums[right] 之间，则在右半部分查找，否则在左半部分查找。
终止条件：
每次根据判断结果缩小搜索区间，直到 left 超过 right，或者找到目标值。
*/

func searchArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断哪一半是有序的
		if nums[left] <= nums[mid] { // 左半部分有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
