package DynamicProgramming

/*乘积最大子数组
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组
（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
测试用例的答案是一个 32-位 整数。
*/
/*
算法思路:动态规划
这里因为要区分正负,所以要两个数组进行维护,一个是最小,一个是最大
fMax[i] = max(nums[i], fMax[i-1]*nums[i], fMin[i-1]*nums[i])
fMin[i] = min(nums[i], fMax[i-1]*nums[i], fMin[i-1]*nums[i])
*/

func maxProduct(nums []int) int {
	n := len(nums)
	fMax := make([]int, n)
	fMin := make([]int, n)
	fMax[0], fMin[0] = nums[0], nums[0]
	maxNum := nums[0]
	for i := 1; i < n; i++ {
		fMax[i] = max(nums[i], fMax[i-1]*nums[i], fMin[i-1]*nums[i])
		fMin[i] = min(nums[i], fMax[i-1]*nums[i], fMin[i-1]*nums[i])
		if fMax[i] > maxNum {
			maxNum = fMax[i]
		}
	}

	return maxNum
}

// 节约空间的做法
func maxProduct1(nums []int) int {
	n := len(nums)
	fMax, fMin, maxNum := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {

		tmpMin := min(nums[i], fMax*nums[i], fMin*nums[i])
		fMax = max(nums[i], fMax*nums[i], fMin*nums[i])
		fMin = tmpMin

		if fMax > maxNum {
			maxNum = fMax
		}
	}
	return maxNum
}
