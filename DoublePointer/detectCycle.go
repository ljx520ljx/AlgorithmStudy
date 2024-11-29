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
