package DFS

// 104. 二叉树的最大深度
/*
算法思路:自顶而下
遍历整个树，记录每个节点的深度，然后维护最大值
*/
func maxDepth(root *TreeNode) int {
	var res int
	var dfs func(node *TreeNode, cnt int)
	dfs = func(node *TreeNode, cnt int) {
		if node == nil {
			return
		}
		res = max(res, cnt)
		dfs(node.Left, cnt+1)
		dfs(node.Right, cnt+1)
	}
	dfs(root, 1)
	return res
}
