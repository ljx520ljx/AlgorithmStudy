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
	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m
	halfLen := (m + n + 1) / 2

	for low <= high {
		// i 是 nums1 的分割点
		i := (low + high) / 2
		// j 是 nums2 的分割点
		j := halfLen - i

		// 边界值
		nums1LeftMax := math.Inf(-1)
		if i > 0 {
			nums1LeftMax = float64(nums1[i-1])
		}
		nums1RightMin := math.Inf(1)
		if i < m {
			nums1RightMin = float64(nums1[i])
		}

		nums2LeftMax := math.Inf(-1)
		if j > 0 {
			nums2LeftMax = float64(nums2[j-1])
		}
		nums2RightMin := math.Inf(1)
		if j < n {
			nums2RightMin = float64(nums2[j])
		}

		// 检查是否满足分割条件
		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			// 中位数计算
			if (m+n)%2 == 0 {
				return (math.Max(nums1LeftMax, nums2LeftMax) + math.Min(nums1RightMin, nums2RightMin)) / 2.0
			}
			return math.Max(nums1LeftMax, nums2LeftMax)
		} else if nums1LeftMax > nums2RightMin {
			// i 太大了，向左调整
			high = i - 1
		} else {
			// i 太小了，向右调整
			low = i + 1
		}
	}

	return 0.0 // 不应该到达这里
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
