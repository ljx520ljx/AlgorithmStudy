package DynamicProgramming

// TreeNode 337. 打家劫舍 III
/*
算法思路:动态规划+递归
1.考虑每个节点的两种状态：
	1.选: 那么两个孩子就不能选了，值为两个孩子不选的最大值与当前节点值之和
	2.不选: 那么两个孩子可以选也可以不选，值为两个孩子选与不选的最大值之和
2.递归边界:
	1.没有节点，怎么选都是0
3.递归过程:
	1.递归左子树
	2.递归右子树
	3.计算当前节点选或不选的最大值
	4.返回当前节点选或不选的最大值
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob2(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		Lrob, Lnotrob := dfs(node.Left)                   // 递归左子树
		Rrob, Rnotrob := dfs(node.Right)                  // 递归右子树
		rob := node.Val + Lnotrob + Rnotrob               // 选
		notrob := max(Lrob, Lnotrob) + max(Rrob, Rnotrob) // 不选
		return rob, notrob
	}
	return max(dfs(root))
}
