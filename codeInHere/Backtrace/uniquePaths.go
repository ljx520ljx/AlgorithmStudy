package Backtrace

//不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径？
/*
算法思路:回溯+动态规划
动态规划表 dp:
dp[i][j] 表示从 (i, j) 点到右下角 (m-1, n-1) 的路径数量。
初始化一个二维切片 dp，大小为 m x n，其所有元素都初始化为 0。
递归回溯 backtrace(i, j):
backtrace(i, j) 是一个递归函数，它返回从 (i, j) 到目标点 (m-1, n-1) 的路径数量。
边界条件：当 i == m-1 或 j == n-1 时，表示到达了网格的最右边或最下边。此时有且只有一种方式到达目标点，即只能一直往右或一直往下走，所以返回 1。
已计算的值：如果 dp[i][j] 已经计算过（即不为 0），说明从 (i, j) 到目标点的路径数量已经求得过，直接返回该值，避免重复计算。
递归计算：对于每一个位置 (i, j)，可以通过两条路径走到目标点：向下走 (i+1, j) 或向右走 (i, j+1)。
所以 dp[i][j] 的值为这两条路径的路径数之和:dp[i][j] = backtrace(i+1, j) + backtrace(i, j+1)
时间复杂度是 O(m * n)，即网格的大小。每个状态 (i, j) 只会被计算一次，并存储在 dp 数组中。
*/

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var backtrace func(int, int) int
	backtrace = func(i, j int) int {
		if i == m-1 || j == n-1 {
			return 1
		}

		if dp[i][j] != 0 {
			return dp[i][j]
		}

		dp[i][j] = backtrace(i+1, j) + backtrace(i, j+1)

		return dp[i][j]
	}

	return backtrace(0, 0)
}
