package DoublePointer

//82. 删除排序链表中的重复元素 II
//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
/*
算法思路:双指针
1. 定义一个虚拟头节点dummy，指向head
2. 定义两个指针pre和curr，分别指向dummy和head.
   pre指向当前不重复的节点，curr指向当前节点
3. 遍历链表，如果curr和curr.Next的值相等，则记录值val，然后循环遍历，直到curr的值不等于val
4. 将pre.Next指向curr.Next，跳过所有重复的节点
*/
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre, curr := dummy, head

	for curr != nil {
		if curr.Next != nil && curr.Next.Val == curr.Val {
			val := curr.Val
			for curr != nil && curr.Val == val {
				curr = curr.Next
			}
			pre.Next = curr
		} else {
			pre = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}
