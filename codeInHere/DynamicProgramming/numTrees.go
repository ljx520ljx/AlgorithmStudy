package DynamicProgramming

// 96. 不同的二叉搜索树
/*
算法思路:动态规划
dp[i] 表示有 i 个节点的二叉搜索树的数量
状态转移方程:dp[i] = dp[j] * dp[i-j-1]
这里的dp[j]表示左子树的节点数量，dp[i-j-1]表示右子树的节点数量
它们的乘积表示以j+1为根节点的二叉搜索树的数量
*/

func numTrees(n int) int {
	// dp[i] 表示有 i 个节点的二叉搜索树的数量
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
