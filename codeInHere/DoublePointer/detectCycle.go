package DoublePointer

//环形链表,判断链表是否有环
/*
算法思路:快慢指针或者hash表
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	nodeMap := make(map[*ListNode]int)
	curr := head
	for curr != nil {
		nodeMap[curr] = 0
		if _, ok := nodeMap[curr.Next]; ok {
			return curr.Next
		}
		curr = curr.Next
	}
	return nil
}

// 快慢指针,自己画图设未知数可解
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for head != slow {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}
