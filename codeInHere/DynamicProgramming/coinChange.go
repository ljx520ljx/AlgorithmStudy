package DynamicProgramming

import "slices"

//零钱兑换整数数组 coins ，表示不同面额的硬币；整数 amount ，表示总金额。
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//每种硬币的数量是无限的。
/*
算法思路:动态规划
dp[i] 表示构成金额 i 所需要的最少硬币数.目标是求出 dp[amount]。
状态转移方程
dp[0] = 0，因为组成金额0不需要硬币。
对于其他金额 i，我们可以通过选择一个硬币 coin 来更新 dp[i]：
dp[i] = min(dp[i], dp[i - coin] + 1)
其中 dp[i - coin] 表示剩余金额 i - coin 所需的最少硬币数，
加上当前硬币 coin 就是总金额 i 所需的硬币数。
*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//给每个dp[]初始化一个大值表示无法构成该金额
	for i := range dp {
		dp[i] = amount + 1
	}
	//排序coins,后面遍历减少重复计算
	slices.Sort(coins)
	dp[0] = 0
	//遍历硬币组
	for _, x := range coins {
		//从x开始,对于那些金额小于硬币面值的就不用遍历节约时间
		for c := x; c <= amount; c++ {
			dp[c] = min(dp[c], dp[c-x]+1)
			//   dp[c] ：不使用当前硬币x的最少硬币数
			//dp[c-x]+1：使用当前硬币x的最少硬币数（就是金额c-x的最优解+1）
		}
	}
	ans := dp[amount]
	//表示能用coins构成所给金额
	if ans < amount+1 {
		return ans
	}
	return -1
}
