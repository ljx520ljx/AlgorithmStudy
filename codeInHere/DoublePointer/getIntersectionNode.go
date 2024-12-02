package DoublePointer

//相交链表:给你两个单链表的头节点headA和headB.
//请你找出并返回两个单链表相交的起始节点.如果两个链表不存在相交节点,返回 null.
/*
算法思路:
问题分析:
两个单链表可能在某一节点后共享相同的部分，这意味着两个链表在某处相交。
如果两个链表不相交，则返回 null。
方法描述:
使用两个指针 pA 和 pB，分别指向链表 A 和链表 B 的头部。
两个指针同时开始遍历各自的链表。
如果某个指针到达链表尾部，则跳转到另一个链表的头部继续遍历。
当两个指针相遇时，它们的相交点即为结果。如果无相交点，则最终两个指针都会指向 nil。
关键点:
当链表有不同长度时，通过切换到另一个链表的头部来平衡差异。
时间复杂度为𝑂(n),空间复杂度为𝑂(1)
边界条件:
如果其中一个链表为空，则直接返回 null。
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}
		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pa
}
