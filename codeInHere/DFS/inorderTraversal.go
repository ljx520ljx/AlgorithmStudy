package DFS

//94. 二叉树的中序遍历
/*
算法思路:递归
先递归处理左子树，然后将根节点的值加入结果集，最后递归处理右子树
*/
func inorderTraversal(root *TreeNode) (res []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}
