package DFS

// LCR 143. 子结构判断 (剑指 Offer 26. 树的子结构)
// 给定两棵二叉树 tree1 和 tree2，判断 tree2 是否是以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
// 注意，空树不会是以 tree1 的某个节点为根的子树具有相同的结构和节点值。
/*
算法思路: 递归
1. 主函数 `isSubStructure(tree1, tree2)`:
   - 目标：判断 tree2 是否是 tree1 的子结构。
   - 边界条件：如果 tree1 为空或 tree2 为空，根据题意（空树不是子结构），直接返回 false。
   - 递归逻辑：
     a. 调用辅助函数 `doesTree1HaveTree2(tree1, tree2)`，检查以 tree1 根节点开始的子树是否与 tree2 完全匹配。
     b. 如果 a 返回 true，则找到了子结构，直接返回 true。
     c. 如果 a 返回 false，则递归地在 tree1 的左子树中寻找 tree2：`isSubStructure(tree1.Left, tree2)`。
     d. 如果 c 返回 true，则找到了子结构，返回 true。
     e. 如果 c 返回 false，则递归地在 tree1 的右子树中寻找 tree2：`isSubStructure(tree1.Right, tree2)`。
     f. 如果 e 返回 true，则找到了子结构，返回 true。
     g. 如果以上所有情况都未找到匹配，则返回 false。
     - 简化的递归表达：`return doesTree1HaveTree2(tree1, tree2) || isSubStructure(tree1.Left, tree2) || isSubStructure(tree1.Right, tree2)`

2. 辅助函数 `doesTree1HaveTree2(node1, node2)`:
   - 目标：判断以 node1 为根的子树是否**包含**以 node2 为根的树（具有相同的结构和节点值）。注意这里的包含是从根节点开始匹配。
   - 递归逻辑和边界条件：
     a. 如果 node2 为空 (`node2 == nil`)：表示 tree2 的这部分已经匹配完毕（或者 tree2 本身就是空的，虽然主函数排除了，但这里要处理），返回 true。
     b. 如果 node1 为空 (`node1 == nil`)：表示 tree1 已经遍历完了，但 tree2 还有节点需要匹配，说明不匹配，返回 false。
     c. 如果 node1 和 node2 的值不相等 (`node1.Val != node2.Val`)：当前节点值不匹配，直接返回 false。
     d. 如果值相等，则需要递归地比较它们的左右子树：
        - 必须同时满足左子树匹配 `doesTree1HaveTree2(node1.Left, node2.Left)` 且右子树匹配 `doesTree1HaveTree2(node1.Right, node2.Right)`。
        - 返回 `doesTree1HaveTree2(node1.Left, node2.Left) && doesTree1HaveTree2(node1.Right, node2.Right)` 的结果。
*/

// isSubStructure 主函数：判断 tree2 是否是 tree1 的子结构
func isSubStructure(tree1 *TreeNode, tree2 *TreeNode) bool {
	// 边界条件：tree1 或 tree2 为空时，不构成子结构关系
	if tree1 == nil || tree2 == nil {
		return false
	}

	// 检查：
	// 1. 从 tree1 的根节点开始，是否能匹配 tree2
	// 2. 或者，tree2 是否是 tree1 左子树的子结构
	// 3. 或者，tree2 是否是 tree1 右子树的子结构
	// 满足任意一个条件即可
	return doesTree1HaveTree2(tree1, tree2) || isSubStructure(tree1.Left, tree2) || isSubStructure(tree1.Right, tree2)
}

// doesTree1HaveTree2 辅助函数：判断以 node1 为根的子树是否包含 node2 (结构和值相同)
func doesTree1HaveTree2(node1 *TreeNode, node2 *TreeNode) bool {
	// 递归终止条件 1: tree2 已匹配完成 (到达 nil 节点)
	if node2 == nil {
		return true // 说明 tree2 这条路径走完了，匹配成功
	}
	// 递归终止条件 2: tree1 遍历完了，但 tree2 还没完，说明不匹配
	if node1 == nil {
		return false
	}
	// 递归终止条件 3: 当前节点值不匹配
	if node1.Val != node2.Val {
		return false
	}

	// 当前节点值匹配，继续递归比较左右子树
	// 必须左右子树都匹配才算成功
	return doesTree1HaveTree2(node1.Left, node2.Left) && doesTree1HaveTree2(node1.Right, node2.Right)
}
