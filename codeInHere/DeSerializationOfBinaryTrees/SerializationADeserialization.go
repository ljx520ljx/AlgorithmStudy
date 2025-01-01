package DeSerializationOfBinaryTrees

import (
	"strconv"
	"strings"
)

//297. 二叉树的序列化与反序列化
/*
算法思路:dfs
Codec结构体作用是把序列化和反序列化这两个方法组合在一起
Constructor() 初始化 Codec 对象
serialize(root) 利用dfs遍历二叉树，把二叉树序列化为字符串
deserialize(data) 把字符串反序列化为二叉树
*/

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

// Constructor 初始化 Codec 对象
func Constructor() Codec {
	return Codec{}
}

// serialize 序列化二叉树为字符串
func (this *Codec) serialize(root *TreeNode) string {
	var result []string
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			result = append(result, "null")
			return
		}
		result = append(result, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	// 把切片用逗号连接起来
	return strings.Join(result, ",")
}

// deserialize 反序列化字符串为二叉树
func (this *Codec) deserialize(data string) *TreeNode {
	// 把字符串用逗号分割成切片
	values := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if len(values) == 0 {
			return nil
		}
		v0 := values[0]
		values = values[1:]
		if v0 == "null" {
			return nil
		}
		num, _ := strconv.Atoi(v0)
		// 递归构建二叉树
		node := &TreeNode{Val: num}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}
