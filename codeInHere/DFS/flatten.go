package DFS

// 114. 二叉树展开为链表
/*
算法思路:头插法
采用头插法构建链表，也就是从节点 6 开始，在 6 的前面插入 5，在 5 的前面插入 4，依此类推。
为此，要按照 6→5→4→3→2→1 的顺序访问节点，也就是按照右子树 - 左子树 - 根的顺序 DFS 这棵树。
DFS 的同时，记录当前链表的头节点为 prev。一开始 prev 是空节点。
*/
func flatten(root *TreeNode) {
	var prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 先递归处理右子树
		dfs(node.Right)
		// 再递归处理左子树
		dfs(node.Left)
		// 将当前节点的右子树连接到 prev 的右子树
		node.Right = prev
		// 将左子树置为 nil
		node.Left = nil
		// 更新 prev 指针
		prev = node
	}
	dfs(root)
}
