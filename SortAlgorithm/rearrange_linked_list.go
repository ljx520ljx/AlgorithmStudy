package SortAlgorithm

//重排链表
/*算法思路:
找到链表的中点：
使用快慢指针的方法，slow 和 fast，来找到链表的中间节点。
这样可以将链表分为前后两部分。
反转链表的后半部分：
从中间节点的下一个节点开始反转后半部分。
我们通过 reverseList 函数将后半部分链表逆序。
合并链表的前后两部分：
合并前半部分和反转后的后半部分，交替连接节点。
使用 mergeLists 函数进行合并。
*/
//定义链表节点
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 主函数，用于重新排列链表
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 步骤1：通过快慢指针找到链表中点
	mid := findMiddle(head)

	// 步骤2：反转后半部分链表
	l2 := reverseList(mid.Next)
	mid.Next = nil // 切断前半部分和后半部分的连接

	// 步骤3：合并前后两部分链表
	mergeLists1(head, l2)
}

// 找到链表的中点，使用快慢指针
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

// 合并两个链表
func mergeLists1(l1, l2 *ListNode) {
	var next1, next2 *ListNode
	for l1 != nil && l2 != nil {
		next1 = l1.Next
		next2 = l2.Next

		l1.Next = l2
		if next1 == nil {
			break
		}
		l2.Next = next1

		l1 = next1
		l2 = next2
	}
}

// 理解解法后自己重写的
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reorderList(head *ListNode) {
//	mid := findmid(head)
//	l2 := reverselist(mid.Next)
//	mid.Next = nil
//	mergelist(head, l2)
//}
//
//func findmid(head *ListNode) *ListNode {
//	slow, fast := head, head
//	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return slow
//}
//
//func reverselist(head *ListNode) *ListNode {
//	var pre *ListNode
//	curr := head
//	for curr != nil {
//		next := curr.Next
//		curr.Next = pre
//		pre = curr
//		curr = next
//	}
//	return pre
//}
//
//func mergelist(l1 *ListNode, l2 *ListNode) {
//
//	for l1 != nil && l2 != nil {
//		next1 := l1.Next
//		next2 := l2.Next
//
//		l1.Next = l2
//		if next1 != nil {
//			l2.Next = next1
//		}
//		l1 = next1
//		l2 = next2
//	}
//}
