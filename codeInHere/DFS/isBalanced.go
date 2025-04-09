package DFS

// TreeNode Definition for a binary tree node.

//110. 平衡二叉树
//给定一个二叉树，判断它是否是高度平衡的。
//对于本题中，一棵高度平衡二叉树定义为：一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。
/*
算法思路: 后序遍历 (Bottom-Up DFS)
1. 核心思想：在计算一个节点高度的同时，判断其子树是否平衡。如果子树不平衡，则无需继续计算，直接返回一个特殊标记（如-1）。
2. 定义递归函数 `getHeight(node)`：
   - 功能：返回以 `node` 为根的子树的高度。如果该子树不平衡，则返回 -1。
   - 基本情况：如果 `node` 为 nil，表示空树，高度为 0（这里我们定义空树高度为0，叶子节点高度为1），是平衡的。
   - 递归步骤：
     a. 递归计算左子树的高度 `leftHeight = getHeight(node.Left)`。
     b. 如果 `leftHeight == -1`，说明左子树已经不平衡，直接返回 -1。
     c. 递归计算右子树的高度 `rightHeight = getHeight(node.Right)`。
     d. 如果 `rightHeight == -1`，说明右子树已经不平衡，直接返回 -1。
     e. 判断当前节点是否平衡：计算左右子树的高度差 `abs(leftHeight - rightHeight)`。如果差值大于 1，说明当前节点不平衡，返回 -1。
     f. 如果当前节点平衡，则返回其高度：`1 + max(leftHeight, rightHeight)`。
3. 主函数 `isBalanced(root)`：
   - 调用 `getHeight(root)`。
   - 如果返回值为 -1，则树不平衡，返回 false。
   - 如果返回值大于等于 0，则树平衡，返回 true。
*/

func isBalanced(root *TreeNode) bool {
	// 调用辅助函数，如果返回 -1 说明不平衡
	return getHeight(root) != -1
}

// getHeight 辅助函数：计算节点高度，如果不平衡则返回 -1
func getHeight(node *TreeNode) int {
	// 基本情况：空节点高度为 0
	if node == nil {
		return 0
	}

	// 递归计算左子树高度
	leftHeight := getHeight(node.Left)
	// 如果左子树不平衡，或在计算中发现不平衡，则提前返回 -1
	if leftHeight == -1 {
		return -1
	}

	// 递归计算右子树高度
	rightHeight := getHeight(node.Right)
	// 如果右子树不平衡，或在计算中发现不平衡，则提前返回 -1
	if rightHeight == -1 {
		return -1
	}

	// 判断当前节点是否平衡（左右子树高度差不超过 1）
	if abs(leftHeight-rightHeight) > 1 {
		// 当前节点不平衡，返回 -1
		return -1
	}

	// 当前节点平衡，返回其高度（左右子树最大高度 + 1）
	return 1 + max(leftHeight, rightHeight)
}

// abs 计算整数绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
