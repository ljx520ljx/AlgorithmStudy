package SlidingWindow

// 209. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target，找出该数组中满足其和 ≥ target 的长度最小的连续子数组，并返回其长度。
// 如果不存在符合条件的子数组，返回 0。
/*
算法思路: 滑动窗口
1. 初始化左右指针 left = 0, right = 0，最小长度 minLen = n+1，当前窗口和 sum = 0
2. 扩大窗口：
   - 移动右指针 right，将新元素加入窗口，更新 sum += nums[right]
3. 缩小窗口：
   - 当 sum >= target 时，说明当前窗口满足条件
   - 记录当前窗口长度 right - left + 1，并更新最小长度
   - 尝试缩小窗口：移动左指针 left，更新 sum -= nums[left]
   - 重复这个过程，直到 sum < target
4. 继续扩大窗口，重复步骤 2 和 3，直到右指针到达数组末尾
5. 最后，如果 minLen 仍为初始值，表示没有找到符合条件的子数组，返回 0；否则返回 minLen
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	// 初始化最小长度为 n+1，任何长度都不可能超过 n
	minLen := n + 1
	// 初始化窗口左边界和当前和
	left, sum := 0, 0

	// 右指针从0开始，向右扩展窗口
	for right := 0; right < n; right++ {
		// 将当前元素加入窗口
		sum += nums[right]

		// 当窗口内元素和大于等于目标值时，尝试缩小窗口
		for sum >= target {
			// 更新最小长度
			minLen = min(minLen, right-left+1)
			// 缩小窗口：移除窗口最左边的元素
			sum -= nums[left]
			left++
		}
	}

	// 如果 minLen 仍为初始值，表示没有找到符合条件的子数组
	if minLen == n+1 {
		return 0
	}

	return minLen
}
