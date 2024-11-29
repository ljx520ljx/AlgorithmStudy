package DoublePointer

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
/*
算法思路:该题的算法思路非常多样
1.遍历两遍,第一遍记录链表长度,第二遍删除节点
2.快慢双指针遍历一遍即可删除节点
3.利用hash表,创建map[int]*ListNode,一次遍历即可
4.利用栈弹出倒数第n个节点然后删除
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

// 栈
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// 计算链表长度,两次遍历
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// Hash表
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	// 创建哑节点并将其指向原链表的头节点
	dummy := &ListNode{Next: head}

	// 使用 curr 从哑节点开始遍历
	curr := dummy
	nodeMap := make(map[int]*ListNode)
	i := 0

	// 遍历链表，将每个节点保存在 nodeMap 中
	for curr != nil {
		nodeMap[i] = curr
		curr = curr.Next
		i++
	}

	// 现在通过 nodeMap 访问要删除节点的前一个节点
	node := nodeMap[i-n-1]
	// 删除目标节点
	node.Next = node.Next.Next

	// 返回哑节点的 Next，即新的头节点
	return dummy.Next
}
