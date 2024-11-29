package DoublePointer

//接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图
//计算按此排列的柱子，下雨之后能接多少雨水。
/*
算法思路:双指针+中间收缩
使用双指针：
设定两个指针 left 和 right，分别指向数组的左右两端。
用 left_max 和 right_max 分别记录从左边和右边遍历时遇到的最大高度。
逐步向中间收缩：
如果 height[left] 小于 height[right]：如果 height[left] 大于 left_max，则更新 left_max。
否则，说明 height[left] 可以接住 left_max - height[left] 单位的水。
left 向右移动。
否则：如果 height[right] 大于 right_max，则更新 right_max。
否则，说明 height[right] 可以接住 right_max - height[right] 单位的水。
right 向左移动。
结束条件：
当 left 和 right 相遇时，算法结束。
*/

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}
