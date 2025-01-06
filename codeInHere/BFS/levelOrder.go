package BFS

// 102. 二叉树的层序遍历
/*
算法思路:bfs
用队列存储每一层的节点，然后遍历每一层的节点，将节点的值存入数组，然后将节点的左右子节点存入队列
*/

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		vals := make([]int, len(q))
		for i := range vals {
			node := q[0]
			q = q[1:]
			vals[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, vals)
	}
	return res
}
