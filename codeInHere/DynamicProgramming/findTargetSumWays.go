package DynamicProgramming

//目标和:给你一个非负整数数组 nums 和一个整数 target 。
//向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
/*
算法思路:动态规划
问题转化：
给定数组 nums 和目标值 target，问题可以转化为从 nums 中选择若干元素，
使得它们的和为 P = (target + sum(nums)) / 2，其中 sum(nums) 是数组所有元素的总和。
这个转化依赖于目标和公式：P - N = target，其中 P 和 N 分别表示加上正号和负号的元素之和。
可行性判断：
如果 target 的绝对值大于 sum(nums)，或者 (target + sum(nums)) 不是偶数，直接返回 0，表示没有解。
动态规划：
初始化 f[0] = 1，因为和为 0 的方式只有 1 种（不选任何元素）。
对于数组中的每个元素 x，从 f[m] 到 f[x] 反向更新，使用上述的状态转移方程 f[i] = f[i] + f[i - x]。
最终，f[m] 即为能够组成和为目标值 P = (target + sum(nums)) / 2 的子集数目。
*/

func findTargetSumWays(nums []int, target int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	s -= abs(target)
	if s < 0 || s%2 == 1 {
		return 0
	}
	m := s / 2

	f := make([]int, m+1)
	f[0] = 1
	for _, x := range nums {
		for c := m; c >= x; c-- {
			f[c] += f[c-x]
		}
	}
	return f[m]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findTargetSumWays1(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}
