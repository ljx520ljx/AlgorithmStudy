package Queue

//二叉树的锯齿形层序遍历
//给你二叉树的根节点 root,返回其节点值的锯齿形层序遍历
//(即先从左往右,再从右往左进行下一层遍历,以此类推,层与层之间交替进行).
/*
算法思路:
层次遍历：
使用队列（Queue）进行层次遍历，每次遍历一整层节点。
在Go中，可以用一个切片来模拟队列。
锯齿形遍历：
使用一个布尔变量 leftToRight 来控制当前层的遍历方向。
如果 leftToRight 为 true，将节点值从左到右加入结果数组。
如果 leftToRight 为 false，则将节点值从右到左加入结果数组。
每遍历完一层，切换 leftToRight 的值。
数据结构：
队列：用于层次遍历，每次一层一层地访问节点。
时间复杂度：
每个节点仅被访问一次，时间复杂度为 O(n)，其中n是节点总数。
空间复杂度为 O(n),因为使用了队列存储节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	isleft := true
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelsize := len(queue)
		level := make([]int, levelsize)
		for i := 0; i < levelsize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isleft {
				level[i] = node.Val
			} else {
				level[levelsize-1-i] = node.Val
			}
		}
		isleft = !isleft
		result = append(result, level)
	}
	return result
}
