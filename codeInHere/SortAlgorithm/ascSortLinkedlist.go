package SortAlgorithm

// ListNode 升序排序链表
/*
算法思路:归并排序
切分链表：
使用快慢指针找到链表的中间节点，slow 和 fast 初始化在 head 上。
fast 每次移动两步，slow 每次移动一步。
当 fast 到达链表末尾时，slow 正好位于链表中间。
这样可以在 O(n) 时间内找到链表的中点。
切断链表，将链表分成两部分，前半部分从 head 开始，后半部分从 mid 开始。
递归排序：
使用 sortList 递归地对前半部分和后半部分进行排序，
直到链表被分割为单个节点为止
（递归基准条件：链表为空或只有一个节点时，直接返回 head）。
合并：
通过 Merge 函数合并两个有序链表。
这个过程是对链表的两个部分依次比较，按顺序将节点链接起来。
Merge 函数使用一个虚拟头节点 dummy，
每次将较小的节点加入到合并链表中，直到其中一个链表为空。
最后将剩余的节点直接加入合并链表。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针找到链表的中间节点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 切断链表
	mid := slow.Next
	slow.Next = nil

	// 递归排序左右两个部分
	left := sortList(head)
	right := sortList(mid)

	// 合并排序后的两个部分
	return merge(left, right)
}

// 合并两个有序链表
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
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
