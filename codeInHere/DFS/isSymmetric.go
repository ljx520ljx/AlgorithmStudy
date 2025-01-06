package DFS

// 101.对称二叉树
/*
算法思路:dfs
先判断传入的两个轴对称节点是否相等
然后递归判断传入的左子树的左节点和右子树的右节点是否相等，左子树的右节点和右子树的左节点是否相等
*/
func isSymmetric(root *TreeNode) bool {
	var dfs func(left, right *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}
