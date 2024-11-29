package ReverseList

//K 个一组翻转链表
//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。

/*算法思路：
双指针检查分组：我们使用 pre 和 end 两个指针。pre 用于记录每组的前一个节点位置，end 用于遍历到每组的最后一个节点，以判断是否有足够的节点进行反转。
分组反转：
找到 k 个节点后，将当前组反转。
反转操作通过 reverseList 函数完成。
反转后，将 pre.Next 指向反转后的新头节点，start.Next 指向下一个未反转的节点。
移动指针：更新 pre 和 end 的位置，进入下一组的反转操作。
终止条件：如果剩余节点少于 k 个，则不再进行反转，保持原顺序。
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// reverseKGroup 反转链表的每 k 个一组节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for {
		// 检查剩余节点是否有 k 个
		end := pre
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}

		// 记录当前组的开始和结束节点
		start := pre.Next
		next := end.Next
		end.Next = nil

		// 反转当前组并重新连接链表
		pre.Next = reverseList(start)
		start.Next = next

		// 移动 pre 到下一组的前一个位置
		pre = start
	}

	return dummy.Next
}

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

//自写
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	dummy := &ListNode{Next: head}
//	pre := dummy
//	for {
//		end := pre
//		for i := 0; i < k && end != nil; i++ {
//			end = end.Next
//		}
//		if end == nil {
//			break
//		}
//		start := pre.Next
//		next := end.Next
//		end.Next = nil
//		pre.Next = reverseList(start)
//		start.Next = next
//		pre = start
//	}
//	return dummy.Next
//}
//
//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	curr := head
//	for curr != nil {
//		nextnode := curr.Next
//		curr.Next = prev
//		prev = curr
//		curr = nextnode
//	}
//	return prev
//}
