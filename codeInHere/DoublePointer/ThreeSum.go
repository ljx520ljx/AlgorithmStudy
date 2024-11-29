package DoublePointer

//三数之和:i!=j!=k,nums[i]+nums[j]+nums[k]==0.请你返回所有和为0且不重复的三元组。
/*
算法思路:排序+双指针
排序数组:
首先对数组进行升序排序，方便使用双指针方法。
排序后，相同的元素会相邻排列，便于去重。
遍历数组，固定一个数 nums[i]：
对每个 nums[i]，将它作为三元组的第一个元素，剩余的两个元素用双指针法从排序后的数组中寻找。
双指针法寻找另两个数：
设置两个指针，left 从 i+1 开始，right 从数组尾部开始。
计算当前三元组的和 sum = nums[i] + nums[left] + nums[right]。
如果 sum == 0，说明找到了一个三元组，将其加入结果集，并跳过重复值。
如果 sum < 0，说明需要更大的值，移动 left 向右。
如果 sum > 0，说明需要更小的值，移动 right 向左。
去重：
对固定的 nums[i]，跳过重复值以避免生成重复的三元组。
对于 left 和 right，在找到一个三元组后，继续移动以跳过相同的值。
时间复杂度：
排序的时间复杂度为O(nlogn)。
总复杂度为O(n^2)。
空间复杂度：
由于只使用了固定大小的额外空间，空间复杂度为
O(1)（不考虑结果存储空间）。
*/

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 跳过重复值
		}

		// 双指针
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 跳过重复值
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
