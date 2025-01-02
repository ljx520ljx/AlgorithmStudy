package DynamicProgramming

// 279. 完全平方数
/*
算法思路:动态规划
dp[i] 表示组成i的完全平方数的最少数量
状态转移方程:dp[i] = min(dp[i-j*j])+1(这里j从1到根号i)
边界条件:dp[0] = 0
*/
import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minNum = min(minNum, dp[i-j*j])
		}
		dp[i] = minNum + 1
	}
	return dp[n]
}
