package DFS

import "math"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//二叉树中的最大路径和
//同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//路径和 是路径中各节点值的总和。
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
/*
算法思路:dfs
考虑两种路径情况：
单边路径：从当前节点出发，沿着左或右子树向下延伸，获取的最大路径和。
经过当前节点的路径：这条路径同时包含左子树、当前节点和右子树。
思路：
对于每个节点，计算其从左子树或右子树传递下来的最大路径和。
计算经过当前节点的最大路径和：left + node + right，其中 left 和 right 是当前节点的左子树和右子树的最大路径和，node 是当前节点的值。
递归返回单边路径和的最大值，即 max(left, right) + node，作为递归返回的路径和。
在遍历过程中维护一个全局最大路径和。
*/

func maxPathSum(root *TreeNode) int {
	// 初始最大路径和为最小值
	maxSum := math.MinInt

	// 辅助函数，返回以当前节点为起点的最大路径和
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子树的最大路径和
		left := max(0, dfs(node.Left))   // 左子树路径和，若为负值，则不贡献
		right := max(0, dfs(node.Right)) // 右子树路径和，若为负值，则不贡献

		// 更新最大路径和
		maxSum = max(maxSum, left+node.Val+right)

		// 返回当前节点的最大单边路径和
		return max(left, right) + node.Val
	}

	// 从根节点开始递归
	dfs(root)

	// 返回全局最大路径和
	return maxSum
}
