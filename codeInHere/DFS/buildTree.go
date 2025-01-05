package DFS

// 105. 从前序与中序遍历序列构造二叉树
/*
算法思路:dfs
先序遍历的第一个元素是根节点，根据这个元素在中序遍历的位置可以将中序遍历分为左右子树
然后一直递归划分，直到先序遍历为空
注意:用hash表时间优化存的index需要位偏移才能对上被我们划分的子树
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	var dfs func(preorder []int, left, right int) *TreeNode
	dfs = func(preorder []int, left, right int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		// 当前根节点值是先序遍历的第一个元素
		node := &TreeNode{Val: preorder[0]}
		// 先找到根节点在中序遍历里的位置再分为左右子树
		index, _ := inorderMap[preorder[0]]
		// 根据划分的左右子树分别遍历
		node.Left = dfs(preorder[1:index-left+1], left, index-1)
		node.Right = dfs(preorder[index-left+1:], index+1, right)
		return node
	}
	return dfs(preorder, 0, len(inorder)-1)
}
