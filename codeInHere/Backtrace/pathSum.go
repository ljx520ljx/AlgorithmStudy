package Backtrace

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 113. 路径总和 II
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
// 叶子节点 是指没有子节点的节点。
/*
算法思路: 深度优先搜索 (DFS) + 回溯 (使用闭包和外部变量)

问题: 找出所有从根节点到叶子节点的路径，使得路径上节点值的总和等于 targetSum。

方法:
1. 外部变量: 在 `pathSum` 函数内部，但在 `dfs` 闭包外部定义：
   - `ans [][]int`: 存储最终所有有效路径的结果列表（使用命名返回值）。
   - `path []int`: 存储当前从根节点到正在访问节点的路径上的节点值。这个切片被 `dfs` 闭包共享和修改。
2. 定义 DFS 闭包 `dfs(node *TreeNode, left int)`:
   - 功能：递归地探索以 `node` 为根的子树，寻找路径和为 `left` 的路径。
   - `node`: 当前访问的节点。
   - `left`: 从根节点到达当前节点后，还需要凑齐的目标和。
   - 基本情况 (Base Case): 如果 `node` 为 `nil`，表示当前路径无效或已到达叶节点的子节点，直接返回。
   - 递归步骤:
     a. 记录路径: 将当前节点 `node.Val` 追加到共享的 `path` 切片中。
     b. 更新剩余和: 从 `left` 中减去当前节点的值 `node.Val`。
     c. 判断叶子节点和满足条件:
        i. 通过 `node.Left == node.Right` 判断当前节点是否是叶子节点（当左右子都为 nil 时它们相等）。
        ii. 如果是叶子节点 **并且** 更新后的 `left` 恰好为 0，说明找到了一条从根到此叶子节点的有效路径。
     d. 保存有效路径: 如果找到有效路径，使用 `slices.Clone(path)` 创建当前 `path` 的一个**深拷贝**，并将这个副本追加到外部的 `ans` 结果列表中。**必须拷贝**，因为 `path` 在回溯时会被修改。
     e. 递归探索子节点:
        i. （条件分支） 如果当前不是满足条件的叶子节点，或者即使是（为了后续的回溯能正确执行），都需要继续探索。
        ii. 递归调用 `dfs(node.Left, left)` 探索左子树。
        iii. 递归调用 `dfs(node.Right, left)` 探索右子树。
     f. 回溯 (恢复现场): 在当前节点 `node` 的左右子树递归调用全部返回之后（即该节点的所有分支已探索完毕），必须将当前节点从 `path` 中移除。通过 `path = path[:len(path)-1]` 实现，使 `path` 恢复到进入该节点之前的状态。
3. 主函数 `pathSum(root *TreeNode, targetSum int)`:
   - 初始化 `ans` 和 `path`。
   - 调用 `dfs(root, targetSum)` 启动递归搜索过程。
   - 返回最终的结果 `ans`。
*/
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	// path: 共享变量，存储当前正在构建的路径
	var path []int

	// dfs: 递归闭包，执行深度优先搜索
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		// 基本情况：如果节点为空，则此路不通，返回
		if node == nil {
			return
		}
		// 处理当前节点：添加到路径，更新剩余目标和
		path = append(path, node.Val) // 记录路径
		sum -= node.Val               // 更新剩余和

		// 检查是否是叶子节点且路径和满足条件
		// node.Left == node.Right 隐含了 node.Left == nil && node.Right == nil
		if node.Left == node.Right && sum == 0 {
			// 找到有效路径，保存其副本到结果中
			ans = append(ans, slices.Clone(path))
			// 注意：即使找到，仍需执行下面的回溯步骤
		} else {
			// 如果不是满足条件的叶子节点，则继续递归探索子节点
			dfs(node.Left, sum)  // 探索左子树
			dfs(node.Right, sum) // 探索右子树
		}

		// 回溯：恢复现场，将当前节点从路径末尾移除
		path = path[:len(path)-1]
	}

	// 从根节点开始启动 DFS
	dfs(root, targetSum)
	// 返回收集到的所有有效路径 (使用命名返回值)
	return ans
}
