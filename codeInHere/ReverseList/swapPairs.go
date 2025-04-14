package ReverseList

// 24. 两两交换链表中的节点
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
// 你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
/*
算法思路: 节点指针操作
1. 关键问题：如何在不修改节点值的情况下，仅通过改变节点指针来交换相邻节点
2. 方法一：迭代法
   - 创建哑节点(dummy)指向链表头部，简化边界情况处理
   - 用指针prev跟踪需要交换的节点对的前一个节点
   - 在每次迭代中：
     a. 获取要交换的两个节点：first = prev.Next, second = prev.Next.Next
     b. 重新连接指针：
        - first.Next = second.Next (使第一个节点指向第二个节点之后的节点)
        - second.Next = first (使第二个节点指向第一个节点)
        - prev.Next = second (使前一个节点指向新的首节点，即原来的第二个节点)
     c. 更新prev为交换后的第二个节点(原first)，准备下一轮交换
3. 方法二：递归法
   - 递归函数定义：交换以head为起点的链表相邻节点，并返回新的头节点
   - 基本情况：如果链表为空或只有一个节点，直接返回head
   - 递归步骤：
     a. 获取新的头节点newHead = head.Next
     b. 递归交换剩余部分：head.Next.Next及之后的节点
     c. 连接交换后的节点：head.Next = 递归结果, newHead.Next = head
     d. 返回新头节点newHead
4. 特殊情况处理：
   - 空链表直接返回nil
   - 只有一个节点的链表直接返回该节点
*/

// 迭代方法实现两两交换
func swapPairs(head *ListNode) *ListNode {
	// 创建哑节点，简化操作
	dummy := &ListNode{Next: head}
	prev := dummy

	// 只要还有成对的节点可以交换
	for prev.Next != nil && prev.Next.Next != nil {
		// 获取要交换的两个节点
		first := prev.Next
		second := first.Next

		// 交换节点
		first.Next = second.Next // 第一个节点指向第二个节点之后的节点
		second.Next = first      // 第二个节点指向第一个节点
		prev.Next = second       // 前一个节点指向新的首节点(原第二个节点)

		// 更新prev指针为交换后的第二个节点(原first)
		prev = first
	}

	return dummy.Next
}

// 递归方法实现两两交换
func swapPairsRecursive(head *ListNode) *ListNode {
	// 基本情况：空链表或只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 保存下一个节点作为新的头节点
	newHead := head.Next

	// 递归交换剩余节点对
	head.Next = swapPairsRecursive(newHead.Next)

	// 连接当前交换的节点对
	newHead.Next = head

	// 返回新的头节点
	return newHead
}
