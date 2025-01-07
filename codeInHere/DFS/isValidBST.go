package DFS

// 98. 验证二叉搜索树
/*
算法思路:递归或者中序遍历
递归:
判断当前节点是否在合法的范围内
然后继续递归判断左右子树是否合法
中序遍历:
将二叉树中序遍历得到一个数组
然后判断数组是否是递增的
如果数组是递增的,则二叉树是二叉搜索树,否则不是二叉搜索树
*/

import "math"

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left, right int) bool
	dfs = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}
		if node.Val <= left || node.Val >= right {
			return false
		}
		return dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

func isValidBST1(root *TreeNode) bool {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}
	return true
}
