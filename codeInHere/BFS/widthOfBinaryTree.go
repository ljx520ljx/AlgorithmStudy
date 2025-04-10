package BFS

// 662. 二叉树最大宽度
// 给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
// 树的 最大宽度 是所有层中最大的 宽度 。
// 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。
// 将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节点，这些 null 节点也计入长度。
/*
算法思路: 层序遍历 (BFS) + 节点位置编号
1. 核心思想：使用 BFS 逐层遍历树。对于每个节点，我们不仅存储节点本身，还存储它在假设的完全二叉树中的位置编号。
2. 位置编号规则：
   - 假设根节点的位置编号为 1。
   - 如果一个节点的位置编号是 `pos`，那么它的左子节点的位置编号是 `2 * pos`，右子节点的位置编号是 `2 * pos + 1`。
3. BFS 过程：
   - 创建一个队列，用于存储待处理的节点及其位置编号。可以定义一个结构体 `QNode{node *TreeNode, pos int}`。
   - 初始化队列，将根节点及其位置编号 1 加入队列。处理根节点为 nil 的情况。
   - 初始化最大宽度 `maxWidth = 0`。
   - 进入 BFS 循环，当队列不为空时：
     a. 获取当前层的节点数量 `levelSize = len(queue)`。
     b. 记录当前层的**起始位置编号** `levelStartPos`，它就是当前队列中第一个元素的位置编号。
     c. 记录当前层的**结束位置编号** `levelEndPos`，在处理完该层所有节点后更新。
     d. 循环 `levelSize` 次，处理当前层的所有节点：
        i. 从队列中取出一个元素 `item = queue[0]`。
        ii. 更新当前层的结束位置 `levelEndPos = item.pos`。
        iii. 如果该节点的左子节点不为空，将其左子节点和对应的位置编号 `2 * item.pos` 加入队列。
        iv. 如果该节点的右子节点不为空，将其右子节点和对应的位置编号 `2 * item.pos + 1` 加入队列。
        v. 将处理过的节点从队列中移除 `queue = queue[1:]`。
     e. **计算当前层的宽度**：`currentWidth = levelEndPos - levelStartPos + 1`。
     f. **更新全局最大宽度**：`maxWidth = max(maxWidth, currentWidth)`。
4. 返回 `maxWidth`。

注意：节点的位置编号可能会变得非常大，特别是在树不平衡且很深的情况下。虽然题目保证最终答案在 32 位整数范围内，但中间的位置编号可能需要使用 64 位整数（`int64`）来防止溢出，或者通过每层重新计算相对编号来避免（但直接使用全局编号更简单）。这里我们先假设标准 `int` 足够存储位置编号。
*/

// QNode 用于存储队列中的节点和其位置编号
type QNode struct {
	node *TreeNode
	pos  int // 使用 1-based 索引
}

func widthOfBinaryTree(root *TreeNode) int {
	// 处理空树情况
	if root == nil {
		return 0
	}

	maxWidth := 0
	// 使用切片模拟队列
	queue := []QNode{{node: root, pos: 1}}

	for len(queue) > 0 {
		levelSize := len(queue)
		// 记录当前层的起始位置 (队列中第一个元素的位置)
		levelStartPos := queue[0].pos
		// levelEndPos 会在循环中被更新为该层最后一个被处理节点的位置
		levelEndPos := levelStartPos

		for i := 0; i < levelSize; i++ {
			// 取出队首元素
			item := queue[0]
			queue = queue[1:] // 出队

			currentNode := item.node
			currentPos := item.pos

			// 更新当前层访问到的最右侧位置
			levelEndPos = currentPos

			// 将非空的子节点及其位置加入队列
			if currentNode.Left != nil {
				queue = append(queue, QNode{node: currentNode.Left, pos: 2 * currentPos})
			}
			if currentNode.Right != nil {
				queue = append(queue, QNode{node: currentNode.Right, pos: 2*currentPos + 1})
			}
		}

		// 计算当前层的宽度
		maxWidth = max(maxWidth, levelEndPos-levelStartPos+1)
	}

	return maxWidth
}
