package LRU

// LRU
/*
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key):如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value):如果关键字 key 已经存在，
则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。
如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
双向链表+哈希表实现
*/
// LRUCache 代表我们缓存的结构
type LRUCache struct {
	capacity int                  // 缓存的容量
	cache    map[int]*DLinkedNode // 使用 map 存储 key 和链表节点的映射
	head     *DLinkedNode         // 虚拟头节点
	tail     *DLinkedNode         // 虚拟尾节点
}

// DLinkedNode 表示双向链表中的一个节点
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// Constructor 使用给定容量初始化 LRU 缓存
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{}, // 虚拟头节点
		tail:     &DLinkedNode{}, // 虚拟尾节点
	}
	lru.head.next = lru.tail // 初始化链表为 head <-> tail
	lru.tail.prev = lru.head
	return lru
}

// Get 如果缓存中存在该键，则检索该键的值
func (lru *LRUCache) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		// 如果存在，则移动该节点到头部（表示最近使用）
		lru.moveToHead(node)
		return node.value
	}
	return -1 // 如果不存在，返回 -1
}

// Put 插入或更新键的值
func (lru *LRUCache) Put(key int, value int) {
	if node, exists := lru.cache[key]; exists {
		// 如果存在，更新节点的值，并将其移动到头部
		node.value = value
		lru.moveToHead(node)
	} else {
		// 如果不存在，创建一个新节点
		newNode := &DLinkedNode{key: key, value: value}
		lru.cache[key] = newNode
		lru.addToHead(newNode)
		if len(lru.cache) > lru.capacity {
			// 超过容量，移除链表尾部的节点（最久未使用）
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
		}
	}
}

// moveToHead 将给定节点移动到列表的头部
func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// addToHead 将一个节点添加到列表的头部（最近使用的）
func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// removeNode 从链表中删除一个节点
func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// removeTail 删除并返回尾节点（lru）
func (lru *LRUCache) removeTail() *DLinkedNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}

//重写
//type LRUCache struct {
//	capacity int
//	cache    map[int]*DlinkedNode
//	head     *DlinkedNode
//	tail     *DlinkedNode
//}
//
//type DlinkedNode struct {
//	key, val   int
//	prev, next *DlinkedNode
//}
//
//func Constructor(capacity int) LRUCache {
//	lru := LRUCache{
//		capacity: capacity,
//		cache:    make(map[int]*DlinkedNode),
//		head:     &DlinkedNode{},
//		tail:     &DlinkedNode{},
//	}
//	lru.head.next = lru.tail
//	lru.tail.prev = lru.head
//	return lru
//}
//
//func (lru *LRUCache) Get(key int) int {
//	if node, ok := lru.cache[key]; ok {
//		lru.moveToHead(node)
//		return node.val
//	}
//	return -1
//}
//
//func (lru *LRUCache) Put(key int, value int) {
//	if node, ok := lru.cache[key]; ok {
//		node.val = value
//		lru.moveToHead(node)
//	} else {
//		newnode := &DlinkedNode{key: key, val: value}
//		lru.cache[key] = newnode
//		lru.addToHead(newnode)
//		if len(lru.cache) > lru.capacity {
//			removed := lru.removeTail()
//			delete(lru.cache, removed.key)
//		}
//	}
//}
//
//func (lru *LRUCache) moveToHead(node *DlinkedNode) {
//	lru.removeNode(node)
//	lru.addToHead(node)
//}
//
//func (lru *LRUCache) removeNode(node *DlinkedNode) {
//	if node.prev != nil {
//		node.prev.next = node.next
//	}
//	if node.next != nil {
//		node.next.prev = node.prev
//	}
//}
//func (lru *LRUCache) addToHead(node *DlinkedNode) {
//	node.prev = lru.head
//	node.next = lru.head.next
//	lru.head.next.prev = node
//	lru.head.next = node
//}
//func (lru *LRUCache) removeTail() *DlinkedNode {
//	node := lru.tail.prev
//	lru.removeNode(node)
//	return node
//}
