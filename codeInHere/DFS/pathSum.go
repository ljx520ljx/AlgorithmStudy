package DFS

//路径总和 III
//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
/*
算法思路:dfs暴力搜索路径
以该节点为起始节点，计算从该节点开始向下的路径节点值之和是否等于 targetSum。
考虑将二叉树中的每个节点都作为起始节点，找出所有满足条件的路径。
节点为空时，终止当前路径的搜索。
当节点值等于当前 target 值时，结果计数加 1。
分别向左、右子节点递归搜索，更新 target 为 target - node.Val，表示继续寻找剩余和的节点。
还有其他情况就是路径不包含当前节点,则当前节点的左右子节点为开头
方法二:前缀和减少重复计算
使用深度优先搜索（DFS）遍历二叉树，同时使用一个映射（cnt）来记录从根节点到当前节点路径上的前缀和出现的次数。
对于每个节点，计算从根节点到该节点的路径和（s），并检查 s - targetSum 的值是否已经在 cnt 映射中。
如果 s - targetSum 在 cnt 中，将其对应的计数累加到结果 ans 中.
cnt[s]--：在完成对左右子节点的搜索后，将当前前缀和 s 的计数减 1，避免对其他分支的计算造成干扰。
初始时，cnt[0] = 1 表示前缀和为 0 的情况出现了 1 次（初始状态，还未开始计算时）。
当遍历到每个节点时，计算前缀和 s，并根据 s - targetSum 的值来查找是否存在满足条件的前缀和。如果存在，说明存在一段路径的和等于 targetSum，将其计数累加到结果中。
对 cnt[s] 进行加 1 和减 1 的操作，能在不同的搜索路径上正确地维护前缀和的计数。
*/

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, target int) {
		// 如果当前节点为空，返回
		if node == nil {
			return
		}
		// 如果当前节点的值等于 target，结果加 1
		if target == node.Val {
			res++
		}
		// 向左子节点递归，更新 target
		dfs(node.Left, target-node.Val)
		// 向右子节点递归，更新 target
		dfs(node.Right, target-node.Val)
	}
	// 从根节点开始递归
	dfs(root, targetSum)
	// 对左子树递归
	if root != nil && root.Left != nil {
		res += pathSum(root.Left, targetSum)
	}
	// 对右子树递归
	if root != nil && root.Right != nil {
		res += pathSum(root.Right, targetSum)
	}
	return res
}

func pathSum1(root *TreeNode, targetSum int) (ans int) {
	cnt := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		s += node.Val
		ans += cnt[s-targetSum]
		cnt[s]++
		dfs(node.Left, s)
		dfs(node.Right, s)
		cnt[s]-- // 恢复现场
	}
	dfs(root, 0)
	return ans
}
