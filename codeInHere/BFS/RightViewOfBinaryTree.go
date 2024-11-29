package BFS

//二叉树的右视图
//给定一个二叉树的 根节点 root，想象自己站在它的右侧，
//按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
/*
算法实现:
广度优先搜索 (BFS)：我们采用层序遍历（BFS），逐层遍历二叉树。
队列数据结构：使用一个队列来存储每一层的节点。对于每一层的节点数量，我们循环遍历所有节点。
右侧视图节点：在每一层中，我们记录最后一个节点的值，即为从右侧看到的节点。该节点会被存储到结果列表中。
遍历过程：对于每个节点，我们先将其左子节点（如果存在）和右子节点（如果存在）加入队列，以保证从左到右遍历每一层。
结果返回：当遍历完所有层后，我们的结果列表中就包含了每一层的右侧视图节点。
*/

// TreeNode 定义了二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rightSideView 函数返回二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var rightmost int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			//删除这个节点,使得下一轮的queue[0]向右前进
			queue = queue[1:]

			// 当层循环遍历到最后,存储的就是当前层的最后一个节点值作为右侧视图节点
			rightmost = node.Val

			// 将左右子节点添加到队列中,广度优先遍历
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 将右侧视图节点值添加到结果集中
		result = append(result, rightmost)
	}

	return result
}

//自写
//func rightSideView(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//	result := []int{}
//	queue := []*TreeNode{root}
//	for len(queue) > 0 {
//		var rightmost int
//		levelSize := len(queue)
//		for i := 0; i < levelSize; i++ {
//			node := queue[0]
//			queue = queue[1:]
//			rightmost = node.Val
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//		}
//		result = append(result, rightmost)
//	}
//	return result
//}
