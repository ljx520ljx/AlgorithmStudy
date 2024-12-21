package DynamicProgramming

//分割等和子集
//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
/*
算法思路:动态规划
首先问题简化,sum为奇数直接返回false,或者数组内最大数打印target=sum/2,也直接返回false
其次动态规划:dp[i][j]表示从数组的 [0,i] 下标范围内选取若干个正整数（可以是 0 个），是否存在一种选取方案使得被选取的正整数的和等于 j。初始时，dp 中的全部元素都是 false。
状态转移方程:
dp[i][j]={dp[i−1][j]=dp[i−1][j−nums[i]]j≥nums[i],dp[i−1][j],j<nums[i]}
空间优化:dp[j]=dp[j] ∣ dp[j−nums[i]]
*/
func canPartition(nums []int) bool {
	sum, maxNum := 0, 0
	n := len(nums)
	for _, v := range nums {
		sum += v
		maxNum = max(maxNum, v)
	}
	target := sum / 2
	if sum%2 == 1 || maxNum > target {
		return false
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		for j := 1; j < target+1; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

// 化二维为一维,因为这一行只和上一行状态有关
func canPartition1(nums []int) bool {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 != 0 {
		return false
	}
	s /= 2 // 注意这里把 s 减半了
	dp := make([]bool, s+1)
	dp[0] = true
	//dp[nums[0]] = true
	s2 := 0
	for _, x := range nums {
		//这里s2+x小于target,就直接从s2+x开始,节约时间
		s2 = min(s2+x, s)
		//倒序遍历,因为如果正序遍历dp[j-x]这里的dp[j-x]就不是上一行的,而是本行已经被更新的了
		for j := s2; j >= x; j-- {
			dp[j] = dp[j] || dp[j-x]
		}
	}
	return dp[s]
}
