package Merge

//合并k个升序链表
//给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中,返回合并后的链表。
/*
算法思路:使用最小堆构建优先队列
使用优先队列：我们将每个链表的第一个节点放入最小堆中，
堆顶元素始终是当前未合并节点中最小的。
迭代构建新链表：从堆中移除最小节点，并将其添加到结果链表中。
然后，将该节点的下一个节点（如果有）加入堆中。
重复该过程：重复从堆中取出最小元素并添加到结果链表，
直到所有节点都被合并到结果链表中。
*/
import (
	"container/heap"
)

// 定义链表节点
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 定义最小堆，保存链表节点
type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 合并 k 个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	// 创建一个虚拟头节点，方便返回结果
	dummy := &ListNode{}
	curr := dummy

	// 初始化最小堆
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// 将每个链表的头节点加入最小堆
	for _, node := range lists {
		if node != nil {
			heap.Push(minHeap, node)
		}
	}

	// 开始从堆中取最小节点，构建合并后的链表
	for minHeap.Len() > 0 {
		// 取出堆顶最小节点
		smallestNode := heap.Pop(minHeap).(*ListNode)
		curr.Next = smallestNode
		curr = curr.Next

		// 如果该节点有下一个节点，将其加入堆中
		if smallestNode.Next != nil {
			heap.Push(minHeap, smallestNode.Next)
		}
	}

	return dummy.Next
}

//type MinHeap []*ListNode
//
//func (h MinHeap) Len() int           { return len(h) }
//func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
//func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *MinHeap) Push(x interface{}) {
//	*h = append(*h, x.(*ListNode))
//}
//
//func (h *MinHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//func mergeKLists(lists []*ListNode) *ListNode {
//	dummy := &ListNode{}
//	curr := dummy
//	minheap := &MinHeap{}
//	heap.Init(minheap)
//	for _, node := range lists {
//		if node != nil {
//			heap.Push(minheap, node)
//		}
//	}
//
//	for minheap.Len() > 0 {
//		smallestnode := heap.Pop(minheap).(*ListNode)
//		curr.Next = smallestnode
//		curr = curr.Next
//		if smallestnode.Next != nil {
//			heap.Push(minheap, smallestnode.Next)
//		}
//	}
//	return dummy.Next
//}
