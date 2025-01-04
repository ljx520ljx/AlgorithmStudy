package DFS

//617. 合并二叉树
/*
算法思路:dfs
当t1为nil时,返回t2
当t2为nil时,返回t1
t1,t2都不为nil时,合并值
之后t1的左子树和t2的左子树合并
接下啦t1的右子树和t2的右子树合并
最后返回t1
*/
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	//当t1为nil时,返回t2
	if t1 == nil {
		return t2
	}
	//当t2为nil时,返回t1
	if t2 == nil {
		return t1
	}
	//t1,t2都不为nil时,合并值
	t1.Val += t2.Val
	//t1的左子树和t2的左子树合并
	t1.Left = mergeTrees(t1.Left, t2.Left)
	//t1的右子树和t2的右子树合并
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

/*
算法思路:dfs
分不同的情况处理
*/
func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var dfs func(node1 *TreeNode, node2 *TreeNode) *TreeNode
	dfs = func(node1 *TreeNode, node2 *TreeNode) *TreeNode {
		switch {
		case node1 != nil && node2 != nil:
			node1.Val = node1.Val + node2.Val
			node1.Left = dfs(node1.Left, node2.Left)
			node1.Right = dfs(node1.Right, node2.Right)
		case node1 == nil && node2 != nil:
			return node2
		case node1 != nil && node2 == nil:
			return node1
		default:
			return nil
		}
		return node1
	}
	return dfs(root1, root2)
}
