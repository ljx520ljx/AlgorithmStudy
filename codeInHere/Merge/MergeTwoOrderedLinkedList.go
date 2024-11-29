package Merge

//合并两个有序链表
/*
算法思路:
新建一个链表来进行合并
从两个链表的最小值进行对比,把Val小的一方的节点加入到新建链表里
直到某一个链表为空,则把剩下的链表接到后面合并完成
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个虚拟头节点
	dummy := &ListNode{}
	current := dummy

	// 比较两个链表的节点并将较小的节点连接到 current 后面
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// 将剩余的链表连接到 current 后面
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	// 返回合并后的链表的头节点
	return dummy.Next
}
