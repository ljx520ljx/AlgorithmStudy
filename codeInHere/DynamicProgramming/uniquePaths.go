package DynamicProgramming

//不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径？
/*
算法思路:动态规划
我们令 dp[i][j] 是到达 i, j 最多路径
状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
对于第一行 dp[0][j]，或者第一列 dp[i][0]，由于都是在边界，所以只能为 1
*/

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 由于 dp[i][j] 仅与第 i 行和第 i−1 行的状态有关，
// 因此我们可以使用滚动数组代替代码中的二维数组，使空间复杂度降低为 O(n)。
func uniquePaths1(m int, n int) int {
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			//这里的d[j-1]是第i行的,加法执行之前的d[j]是第i-1行的,执行之后是第i行的.
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
