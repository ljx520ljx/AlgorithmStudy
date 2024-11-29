package ReverseList

//回文链表
/*
算法思路:
快慢指针找到链表中点;
反转后半部分链表：从慢指针的位置开始，反转链表的后半部分;
比较前后两部分链表：将反转后的后半部分与前半部分进行比较;
如果所有对应节点的值相同，则是回文链表；否则不是。
恢复链表（可选）：为了不改变输入链表的结构，可以选择将后半部分再次反转回原始顺序。
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Step 1: 使用快慢指针找到链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: 反转后半部分链表
	prev := reverseList(slow.Next)
	//切断slow.Next
	slow.Next = nil

	// Step 3: 比较前半部分和反转后的后半部分
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

//func reverseList(head *ListNode) *ListNode {
//	var pre *ListNode
//	curr := head
//	for curr != nil {
//		nextnode := curr.Next
//		curr.Next = pre
//		pre = curr
//		curr = nextnode
//	}
//	return pre
//}
