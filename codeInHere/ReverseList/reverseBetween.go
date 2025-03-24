package ReverseList

// 92.反转链表II.反转一个链表的"中间部分"
/*
算法思路:
方法一:记录left-1和right+1的位置,然后翻转left-right的链表,最后拼接起来.
方法二:找到left的前一个节点和left节点,使用头插法反转.
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}

	// 创建一个虚拟头节点，简化边界条件处理
	dummy := &ListNode{Next: head}
	curr1 := dummy

	// 找到 left 的前一个节点
	for i := 1; i < left; i++ {
		curr1 = curr1.Next
	}
	curr2 := curr1.Next
	for i := left; i < right; i++ {
		node := curr2.Next
		curr2.Next = node.Next
		node.Next = curr1.Next
		curr1.Next = node
	}
	return dummy.Next
}

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}

	// 创建一个虚拟头节点，简化边界条件处理
	dummy := &ListNode{Next: head}
	curr1 := dummy

	// 找到 left 的前一个节点
	for i := 1; i < left; i++ {
		curr1 = curr1.Next
	}

	// curr2 是翻转链表的尾节点
	curr2 := curr1.Next
	for i := left; i < right; i++ {
		curr2 = curr2.Next
	}

	// 储存翻转链表的头节点
	next1 := curr1.Next
	// 存储第二段没被翻转的链表的头节点
	next2 := curr2.Next
	// 切断
	curr2.Next = nil
	// head2 是链表被翻转后的头节点
	head2 := reverseList(next1)
	// 拼接链表
	curr1.Next = head2
	// 翻转后变成尾节点
	next1.Next = next2

	return dummy.Next
}

/*
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
*/
