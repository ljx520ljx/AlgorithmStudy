package DoublePointer

/*除自身以外数组的乘积
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
*/
/*
算法思路:分成左右两个部分,双指针
初始化两个空数组 L 和 R。对于给定索引 i，L[i] 代表的是 i 左侧所有数字的乘积，R[i] 代表的是 i 右侧所有数字的乘积。
我们需要用两个循环来填充 L 和 R 数组的值。对于数组 L，L[0] 应该是 1，因为第一个元素的左边没有元素。对于其他元素：L[i] = L[i-1] * nums[i-1]。
同理，对于数组 R，R[length-1] 应为 1。length 指的是输入数组的大小。其他元素：R[i] = R[i+1] * nums[i+1]。
当 R 和 L 数组填充完成，我们只需要在输入数组上迭代，且索引 i 处的值为：L[i] * R[i]。
*/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n) // 创建一个长度为 n 的数组
	for i := 0; i < n; i++ {
		answer[i] = 1 // 将 answer 数组的所有元素初始化为 1
	}

	left, right := 0, n-1 // 左右指针初始化
	lp, rp := 1, 1        // 左边乘积和右边乘积初始化为 1

	// 当左指针小于右指针时，继续计算
	for right >= 0 && left < n {
		answer[right] *= rp // 更新右边元素的乘积
		answer[left] *= lp  // 更新左边元素的乘积
		lp *= nums[left]    // 更新左边乘积，并移动左指针
		rp *= nums[right]   // 更新右边乘积，并移动右指针
		left++              // 移动左指针
		right--             // 移动右指针
	}

	return answer // 返回最终结果
}

func productExceptSelf1(nums []int) []int {
	length := len(nums)
	L, R, answer := make([]int, length), make([]int, length), make([]int, length)
	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}
	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}

func productExceptSelf2(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}
	return answer
}
