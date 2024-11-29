package SortAlgorithm

// 排序奇升偶降链表
/*
算法思路:
拆分链表：遍历链表，把奇数位置和偶数位置的节点分别拆成两个链表。
反转偶数链表：将偶数位置的链表反转，以便合并时满足升序要求。
合并链表：交替合并奇数链表和反转后的偶数链表。
*/

// 定义链表节点结构
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 主函数，重排序链表
func reorderOEList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1. 拆分链表
	odd, even := splitList(head)

	// 2. 反转偶数链表
	even = reverseList(even)

	// 3. 合并两个有序链表
	return mergeLists(odd, even)
}

// 拆分链表为奇数链表和偶数链表
func splitList(head *ListNode) (*ListNode, *ListNode) {
	odd, even := head, head.Next
	oddHead, evenHead := odd, even

	for odd != nil && odd.Next != nil && even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	return oddHead, evenHead
}

// // 反转链表
//
//	func reverseList(head *ListNode) *ListNode {
//		var prev *ListNode
//		curr := head
//		for curr != nil {
//			next := curr.Next
//			curr.Next = prev
//			prev = curr
//			curr = next
//		}
//		return prev
//	}
//
// 合并两个链表
func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		current.Next = l1
		l1 = l1.Next
		current = current.Next

		current.Next = l2
		l2 = l2.Next
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}
