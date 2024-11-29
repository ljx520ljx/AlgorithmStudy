package DynamicProgramming

//最大子数组之和
/*
算法思路:动态规划
动态规划的状态转移：
设 dp[i] 表示以 nums[i] 结尾的最大子数组和。
递推公式为：
dp[i] = max(dp[i-1] + nums[i], nums[i])
其中，dp[i-1] + nums[i] 表示将当前元素 nums[i] 加入到之前的最大子数组中，nums[i] 则表示从当前位置开始一个新的子数组。
初始化：
dp[0] 为 nums[0]，即数组的第一个元素。
迭代求解：
通过遍历整个数组，计算每个位置的 dp[i]，同时更新 max_sum 记录到目前为止的最大子数组和。
优化空间：
由于我们只需要 dp[i-1] 和 nums[i] 来计算当前的 dp[i]，所以可以将空间优化到 O(1)，只用一个变量来存储 dp[i-1]。
*/

// O(n)时间复杂度
func maxSubArray(nums []int) int {
	// 初始化当前子数组和
	currentSum := nums[0]
	// 初始化最大和
	maxSum := nums[0]

	// 从第二个元素开始遍历
	for i := 1; i < len(nums); i++ {
		// 更新当前子数组和
		currentSum = max(currentSum+nums[i], nums[i])
		// 更新最大和
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

//分治法实现O(nlog n)时间复杂度
//func maxSubArrayHelper(nums []int, left, right int) int {
//	if left == right {
//		return nums[left] // base case: only one element
//	}
//
//	mid := (left + right) / 2
//
//	// recursively calculate the max subarray sum for left and right parts
//	leftMax := maxSubArrayHelper(nums, left, mid)
//	rightMax := maxSubArrayHelper(nums, mid+1, right)
//
//	// calculate the max subarray sum crossing the middle
//	leftCrossMax := nums[mid]
//	rightCrossMax := nums[mid+1]
//	temp := 0
//
//	// Calculate left part maximum sum that crosses the middle
//	for i := mid; i >= left; i-- {
//		temp += nums[i]
//		if temp > leftCrossMax {
//			leftCrossMax = temp
//		}
//	}
//
//	temp = 0
//	// Calculate right part maximum sum that crosses the middle
//	for i := mid + 1; i <= right; i++ {
//		temp += nums[i]
//		if temp > rightCrossMax {
//			rightCrossMax = temp
//		}
//	}
//
//	crossMax := leftCrossMax + rightCrossMax
//
//	// return the maximum of the three values
//	return max(max(leftMax, rightMax), crossMax)
//}
//
//// maxSubArray returns the maximum subarray sum using divide and conquer
//func maxSubArray(nums []int) int {
//	return maxSubArrayHelper(nums, 0, len(nums)-1)
//}
