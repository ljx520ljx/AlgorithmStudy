package ReverseList

//反转链表递归实现

// 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归反转链表
func reverseList1(head *ListNode) *ListNode {
	// 基本条件：如果链表为空或只有一个节点，直接返回head
	if head == nil || head.Next == nil {
		return head
	}

	// 递归反转剩余链表
	newHead := reverseList1(head.Next)

	// 反转当前节点指向
	head.Next.Next = head
	head.Next = nil

	return newHead
}
