package DFS

//二叉树的最近公共祖先
/*
算法思路:使用DFS深度优先搜索
基本条件：
如果当前节点 root 为空，或等于节点 p 或 q，则直接返回当前节点 root。
这表示当前节点是一个可能的公共祖先。
递归查找：
对 root 的左子树和右子树分别进行递归调用，查找 p 和 q 的公共祖先，
结果分别存储在 left 和 right 中。
判断是否是最近公共祖先：
如果 left 和 right 均不为空，说明 p 和 q 分别在当前节点的左子树和右子树中，
因此当前节点 root 就是最近的公共祖先。
返回非空子树：
如果 left 或 right 其中一个为空，则返回非空的那一边.
这样最终会找到包含 p 和 q 的子树的根节点。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 基本条件，如果当前节点为空或等于 p 或 q，则返回当前节点
	if root == nil || root == p || root == q {
		return root
	}

	// 递归查找左子树和右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左子树和右子树分别包含 p 和 q，则当前节点 root 为最近公共祖先
	if left != nil && right != nil {
		return root
	}

	// 否则返回非空的子树
	if left != nil {
		return left
	}
	return right
}
