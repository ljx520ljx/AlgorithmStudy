package DynamicProgramming

func walk(n, k int, data []int) int {
	// dp[i][j] 表示走完第 i+1 天，背包中还剩余 j 食物时所花费的最少钱数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k)
	}
	// 初始化第一天走完,背包还剩多少食物时的花费
	for i := 0; i < k; i++ {
		dp[0][i] = (i + 1) * data[0]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			// 不购买食物，直接继承前一天的状态
			if j > 0 {
				dp[i][j-1] = dp[i-1][j]
			}
			// 考虑购买食物的情况
			for buy := 0; buy <= k-j; buy++ {
				// 计算购买 buy 量食物后的花费
				cost := dp[i-1][j] + buy*data[i]
				// 更新最小花费
				if dp[i][j+buy-1] == 0 || cost < dp[i][j+buy-1] {
					dp[i][j+buy-1] = cost
				}
			}
		}
	}
	return dp[n-1][0]
}
