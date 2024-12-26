package DynamicProgramming

// 312 戳气球
/*
算法思路:动态规划
1.创建一个新的数组,给边界添加哨兵值1
2.dp[start][end] 表示戳破从 start 到 end 区间内戳气球能获得的最大硬币数。
3.状态转移方程:位置k被当作最后一个戳破的气球时,能获得的硬币数为:dp[start][k-1] + dp[k+1][end] + newNums[start-1]*newNums[k]*newNums[end+1]
4.最终结果:dp[1][n]
*/

func maxCoins(nums []int) int {
	n := len(nums)
	// 创建一个新的数组，添加边界的虚拟气球值为 1
	newNums := make([]int, n+2)
	newNums[0] = 1
	newNums[n+1] = 1
	for i := 1; i <= n; i++ {
		newNums[i] = nums[i-1]
	}
	// 创建二维数组用于存储中间结果
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	// 从长度为 1 的子问题开始求解
	for length := 1; length <= n; length++ {
		for start := 1; start <= n-length+1; start++ {
			end := start + length - 1
			maxNum := 0
			for k := start; k <= end; k++ {
				//这里的k表示在start到end区间内最后一个戳破的气球的位置
				coins := dp[start][k-1] + dp[k+1][end] + newNums[start-1]*newNums[k]*newNums[end+1]
				if coins > maxNum {
					maxNum = coins
				}
			}
			dp[start][end] = maxNum
		}
	}
	return dp[1][n]
}
