package DFS

//把二叉搜索树转换为累加树
//把每个节点的值替换为所有节点值大于等于该节点值的节点值总和
/*
算法思路:反向中序遍历
反序中序遍历该二叉搜索树，记录过程中的节点值之和，
并不断更新当前遍历到的节点的节点值，即可得到题目要求的累加树。
*/

func convertBST(root *TreeNode) *TreeNode {
	s := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		s += node.Val
		node.Val = s
		dfs(node.Left)
	}
	dfs(root)
	return root
}
