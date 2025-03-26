package Search

import "math"

//寻找两个有序数组的中位数
/*
算法思路:二分查找
二分查找法：
为了满足(log(m+n)) 的时间复杂度，我们不能简单地合并两个数组后排序，而是采用二分查找的思想。
假设将两个数组分成两个部分，左边部分包含前𝑘小的数，右边部分包含其余的数，中位数正好是左边部分的最大值或两个部分的平均值。
分割逻辑：
使用二分查找较短数组（假设 nums1 总是较短）。
对 nums1 的分割点进行二分，找到一个分割点，使得左部分的最大值小于等于右部分的最小值。
分割条件：
定义切割点i在 nums1 中，𝑗在 nums2 中：i+j=(m+n+1)/2
满足:nums1[i−1]≤nums2[j]且nums2[j−1]≤nums1[i]
根据条件调整i的范围直到找到合适的分割点。
特殊情况处理：
如果一个数组为空，直接返回另一个数组的中位数。
分割点可能触及数组边界，需单独判断。
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保nums1的长度小于等于nums2，这样可以减少边界情况的处理
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	// 二分查找的目标是找到一个位置，将两个数组分成左右两部分，使得：
	// 1. 左半部分的长度等于右半部分（或比右半部分多1个元素）
	// 2. 左半部分的最大值小于等于右半部分的最小值
	for left <= right {
		// partitionX是nums1的分割点
		partitionX := (left + right) / 2
		// partitionY是nums2的分割点，保证左半部分总长度为(m+n+1)/2
		partitionY := (m+n+1)/2 - partitionX

		// 获取分割点左右的值
		var maxLeftX, minRightX, maxLeftY, minRightY int

		// 处理nums1分割点的边界情况
		if partitionX == 0 {
			maxLeftX = math.MinInt32
		} else {
			maxLeftX = nums1[partitionX-1]
		}
		if partitionX == m {
			minRightX = math.MaxInt32
		} else {
			minRightX = nums1[partitionX]
		}

		// 处理nums2分割点的边界情况
		if partitionY == 0 {
			maxLeftY = math.MinInt32
		} else {
			maxLeftY = nums2[partitionY-1]
		}
		if partitionY == n {
			minRightY = math.MaxInt32
		} else {
			minRightY = nums2[partitionY]
		}

		// 检查是否找到正确的分割点
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// 找到了正确的分割点，计算中位数
			if (m+n)%2 == 0 {
				// 偶数个元素，取两部分的平均值
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				// 奇数个元素，取左半部分的最大值
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			// nums1的左半部分太大，需要向左移动
			right = partitionX - 1
		} else {
			// nums1的左半部分太小，需要向右移动
			left = partitionX + 1
		}
	}

	return 0.0 // 不会达到这里
}

//简单合并数组法,时间复杂度为0(m+n)
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	// 合并两个有序数组
//	merged := mergeSortedArrays(nums1, nums2)
//	totalLen := len(merged)
//
//	// 计算中位数
//	if totalLen%2 == 1 {
//		// 奇数长度，返回中间元素
//		return float64(merged[totalLen/2])
//	}
//	// 偶数长度，返回中间两个元素的平均值
//	mid1, mid2 := merged[totalLen/2-1], merged[totalLen/2]
//	return float64(mid1+mid2) / 2.0
//}
//
//func mergeSortedArrays(nums1 []int, nums2 []int) []int {
//	// 初始化结果数组和指针
//	merged := []int{}
//	i, j := 0, 0
//
//	// 合并数组
//	for i < len(nums1) && j < len(nums2) {
//		if nums1[i] < nums2[j] {
//			merged = append(merged, nums1[i])
//			i++
//		} else {
//			merged = append(merged, nums2[j])
//			j++
//		}
//	}
//
//	// 将剩余元素加入结果数组
//	for i < len(nums1) {
//		merged = append(merged, nums1[i])
//		i++
//	}
//	for j < len(nums2) {
//		merged = append(merged, nums2[j])
//		j++
//	}
//
//	return merged
//}
