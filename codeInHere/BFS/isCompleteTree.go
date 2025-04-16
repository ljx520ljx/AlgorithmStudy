package BFS

/*
算法思路: 使用层序遍历 (BFS) 判断二叉树的完全性

完全二叉树的定义:
1. 除了最后一层外，所有层都必须完全填满。
2. 最后一层的所有节点都必须尽可能地集中在左侧。

核心思想:
利用层序遍历的特性，我们可以逐层检查节点。完全二叉树在层序遍历中有一个关键特征：
一旦我们遇到了一个空节点 (nil)，那么在这之后遇到的所有节点都必须是空节点。
如果在一个空节点之后又遇到了非空节点，那么这棵树就不是完全二叉树（因为节点没有尽可能地靠左排列，或者中间出现了“断层”）。

步骤:
1. 初始化:
   - 创建一个队列 (可以用 Go 的切片模拟)。
   - 将根节点 `root` 加入队列。
   - 定义一个布尔标志 `foundNil`，初始为 `false`，用于标记是否已经遇到过 `nil` 节点。
2. 层序遍历 (BFS):
   - 当队列不为空时，循环执行：
     a. 从队列中取出一个节点 `node`。
     b. **检查节点**:
        i. 如果 `node` 是 `nil`:
           - 将 `foundNil` 标记为 `true`。这表示我们开始进入可能出现空缺的阶段。
        ii. 如果 `node` 不是 `nil`:
           - **关键判断**: 检查 `foundNil` 是否已经是 `true`。如果是，说明我们在之前已经遇到了 `nil` 节点，但现在又遇到了一个非 `nil` 节点。这违反了完全二叉树的定义（节点没有靠左，或者有空隙），直接返回 `false`。
           - **入队**: 如果 `foundNil` 仍然是 `false`，则将当前节点 `node` 的左孩子 `node.Left` 和右孩子 `node.Right` **依次** 加入队列。**注意**: 即使孩子是 `nil`，也要将它们加入队列，因为我们需要通过它们来检测后续是否还有非 `nil` 节点。
3. 完成:
   - 如果队列为空，整个遍历过程都没有触发返回 `false` 的条件，说明这棵树满足完全二叉树的定义，返回 `true`。

时间复杂度: O(N)，其中 N 是树中的节点数，因为每个节点最多入队和出队一次。
空间复杂度: O(W)，其中 W 是树的最大宽度。在最坏的情况下（满二叉树），W 约等于 N/2，所以空间复杂度也是 O(N)。
*/

// isCompleteTree 检查二叉树是否为完全二叉树
func isCompleteTree(root *TreeNode) bool {
	// 空树被认为是完全二叉树
	if root == nil {
		return true
	}

	// 使用切片模拟队列进行层序遍历
	queue := []*TreeNode{root}
	// 标志位，标记是否在层序遍历中遇到了 nil 节点
	foundNil := false

	for len(queue) > 0 {
		// 出队
		node := queue[0]
		queue = queue[1:] // 移除队首元素

		if node == nil {
			// 如果遇到 nil，将标志位设为 true
			foundNil = true
		} else {
			// 如果 node 不是 nil

			// 关键检查：如果在遇到 nil 之后又遇到了非 nil 节点，则不是完全二叉树
			if foundNil {
				return false
			}

			// 将左右孩子（即使是 nil）依次加入队列
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// 如果遍历完成没有违反规则，则是完全二叉树
	return true
}
