package DoublePointer

//环形链表,判断链表是否有环
/*
算法思路:快慢指针或者hash表
*/

// hash
func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]int)
	curr := head
	for curr != nil {
		nodeMap[curr] = 0
		if _, ok := nodeMap[curr.Next]; ok {
			return true
		}
		curr = curr.Next
	}
	return false
}

// 快慢指针
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func hasCycle2(head *ListNode) bool {
	slow, fast := head, head // 乌龟和兔子同时从起点出发
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 乌龟走一步
		fast = fast.Next.Next // 兔子走两步
		if fast == slow {     // 兔子追上乌龟（套圈），说明有环
			return true
		}
	}
	return false // 访问到了链表末尾，无环
}
