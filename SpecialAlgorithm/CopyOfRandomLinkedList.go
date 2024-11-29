package SpecialAlgorithm

//随机链表复制
/*
算法思路:
创建复制节点：
对于原链表中的每个节点 n，创建一个新的节点 n'，
并将 n' 插入到 n 的后面。此步骤不会设置 random 指针。
这样我们得到一个“交织”链表，链表中每两个节点是原始节点和复制节点的配对。
设置 random 指针：
对于每个原节点 n，如果 n 的 random 指针指向节点 r，
那么 n' 的 random 指针应该指向 r'。由于 n' 是在 n 的后面，
因此 n.Random.Next 会指向正确的复制节点 r'。
分离链表：
现在，链表已经交织。我们将每个复制节点从交织链表中提取出来，形成新的深拷贝链表。
时间复杂度和空间复杂度：
时间复杂度：O(n)，因为我们对每个节点只遍历了常数次。
空间复杂度：O(1)（除了结果链表外），因为我们是在原地进行修改。
*/
// Node 定义链表节点的结构

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList 实现链表的深拷贝
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 第一步：在每个节点后创建复制节点
	for curr := head; curr != nil; curr = curr.Next.Next {
		copy := &Node{Val: curr.Val}
		copy.Next = curr.Next
		curr.Next = copy
	}

	// 第二步：设置复制节点的 Random 指针
	for curr := head; curr != nil; curr = curr.Next.Next {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
	}

	// 第三步：分离链表
	newHead := head.Next
	for curr := head; curr != nil; curr = curr.Next {
		copy := curr.Next
		curr.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
	}
	return newHead
}
