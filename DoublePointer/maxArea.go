package DoublePointer

//盛最多水的容器
//给定一个长度为 n 的整数数组 height.有n条垂线
//第i条线的两个端点是(i,0)和(i,height[i])
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//返回容器可以储存的最大水量。
/*
算法思路:双指针
问题分析：
给定数组height，每个值代表垂直线的高度。
容器的宽度是两条垂直线的水平距离，容积由较短的垂直线高度和宽度的乘积决定。
目标是找到两条垂直线，使容器可以容纳的水量最大。
双指针法：
使用两个指针分别指向数组的两端（left和right）。
计算当前容器的水量width * min(height[left], height[right])。
更新最大水量maxWater。
移动指针的策略是向中间移动较短的一侧，因为水量受限于较短的垂直线，只有可能通过增加高度获得更大的容积。
复杂度分析：
时间复杂度：O(n)，每次循环移动一个指针，最多遍历一次数组。
空间复杂度：O(1)，只使用了常数额外空间。
*/

func maxArea(height []int) int {
	left, right := 0, len(height)-1 // 初始化左右指针
	maxWater := 0                   // 最大水量

	for left < right {
		// 计算当前容器的水量
		width := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
		} else {
			h = height[right]
		}
		currentWater := width * h
		if currentWater > maxWater {
			maxWater = currentWater
		}

		//这段冗余了,可以加在上面
		// 移动较小的一侧指针
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

// 简洁版
func maxArea1(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		width := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		currWater := width * h
		if currWater > maxWater {
			maxWater = currWater
		}
	}
	return maxWater
}
