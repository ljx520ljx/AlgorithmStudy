package DFS

//LCR174.寻找二叉搜索树中第k大的节点
// 给定一棵二叉搜索树，请找出其中第 k 大的节点。
/*
算法思路:逆序中序遍历
把逆序遍历的结果存储到一个数组中.
当数组的长度等于k时,返回数组的最后一个元素.
*/
func findTargetNode(root *TreeNode, cnt int) int {
	var res []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || len(res) >= cnt {
			return
		}
		dfs(node.Right)
		res = append(res, node.Val)
		dfs(node.Left)
	}
	dfs(root)
	return res[cnt-1]
}
