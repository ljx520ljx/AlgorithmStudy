package DFS

//二叉树的直径 给你一棵二叉树的根节点，返回该树的 直径 。
//二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
//两节点之间路径的 长度 由它们之间边数表示。
/*
算法思路:dfs
递归地计算其左子树和右子树的最大深度（即 lLen 和 rLen），每次递归深度会加 1，表示从该节点到其子节点的路径。
lLen := dfs(node.Left) + 1：表示从当前节点到其左子树的最大深度，+1 是因为我们从当前节点到左子树的路径也算一个边。
rLen := dfs(node.Right) + 1：同理，表示从当前节点到其右子树的最大深度。
对于每个节点，树的直径可能经过该节点，即左右子树的深度之和：lLen + rLen。因此，更新全局变量 ans，记录当前的最大直径
dfs 函数返回当前节点的最大深度（即左右子树中较大的深度加 1）
*/

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lLen := dfs(node.Left) + 1  // 左子树最大链长+1
		rLen := dfs(node.Right) + 1 // 右子树最大链长+1
		ans = max(ans, lLen+rLen)   // 两条链拼成路径
		return max(lLen, rLen)      // 当前子树最大链长
	}
	dfs(root)
	return
}
