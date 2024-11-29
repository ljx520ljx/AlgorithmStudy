package DynamicProgramming

//最大正方形:在一个由 '0' 和 '1' 组成的二维矩阵内
//找到只包含 '1' 的最大正方形，并返回其面积。
/*
算法思路:动态规划
动态规划（DP）方法：定义一个二维数组 dp
其中 dp[i][j] 表示以矩阵中 (i, j) 为右下角的最大正方形的边长。
状态转移：
如果 matrix[i][j] == '1'，则 dp[i][j] 可以由左、上、左上角三个相邻位置的 dp 值计算得到：
dp[i][j]=min(dp[i−1][j],dp[i][j−1],dp[i−1][j−1])+1
这意味着，当前点 (i, j) 能形成的最大正方形的边长是其左、上、左上角相邻的正方形中的最小边长加一。
如果 matrix[i][j] == '0'，则 dp[i][j] = 0，因为不能以该点为正方形的右下角。
边界条件：当矩阵中的 i == 0 或 j == 0 时，dp[i][j] 只能等于 matrix[i][j] 本身，
意味着每个边缘上的元素最多只能组成一个边长为1的正方形。
返回值：最后，遍历整个 dp 数组，找到最大的 dp[i][j] 值，计算其平方即为所求最大正方形的面积。
*/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxlength := 0
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
				maxlength = max(maxlength, dp[i][j])
			}
		}
	}
	return maxlength * maxlength
}
