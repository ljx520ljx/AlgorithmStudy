package LFU

//实现LFU
/*
算法思路:
使用计数最小的键需要被优先移除。
如果有计数相同的键，则移除最近最少使用的键（LRU）。
get 和 put 操作需要平均时间复杂度为 O(1)。
数据结构:
键值对存储（Map）：用来快速查找键对应的值。
频率表（Map of LinkedHashSet）：将使用频率作为键，值是一个有序集合（如 LinkedHashSet），用于存储具有相同频率的键。该集合按最近访问时间排序。
辅助映射（Map）：记录每个键的频率。
最小频率变量：跟踪当前缓存中最小的使用频率。
实现逻辑:
每次 get 操作：找到键值后，将其使用频率加 1，并在频率表中更新位置。
每次 put 操作：
如果键已存在，更新其值，并增加其使用频率。
如果键不存在，则检查容量。如果容量满了，需要删除频率最低的键（最小频率中最久未使用的键），再插入新键。
通过频率表和辅助映射实现 O(1) 的插入和删除操作。
*/

import (
	"container/list"
)

type LFUCache struct {
	capacity   int                // 缓存容量
	minFreq    int                // 当前最小频率
	keyToValue map[int]int        // 键值对
	keyToFreq  map[int]int        // 键到频率的映射
	freqToKeys map[int]*list.List // 频率到键集合的映射
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		minFreq:    0,
		keyToValue: make(map[int]int),
		keyToFreq:  make(map[int]int),
		freqToKeys: make(map[int]*list.List),
	}
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}

	// 如果 key 不存在，返回 -1
	if _, exists := this.keyToValue[key]; !exists {
		return -1
	}

	// 增加 key 的频率
	this.incrementFreq(key)
	return this.keyToValue[key]
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	// 如果 key 已经存在，更新值并增加频率
	if _, exists := this.keyToValue[key]; exists {
		this.keyToValue[key] = value
		this.incrementFreq(key)
		return
	}

	// 如果缓存已满，移除最不常用的键
	if len(this.keyToValue) >= this.capacity {
		this.removeMinFreqKey()
	}

	// 插入新键值对
	this.keyToValue[key] = value
	this.keyToFreq[key] = 1
	if this.freqToKeys[1] == nil {
		this.freqToKeys[1] = list.New()
	}
	this.freqToKeys[1].PushBack(key)
	this.minFreq = 1
}

// 增加某个键的使用频率，并将其在 freqToKeys 映射中从当前频率组移动到下一个频率组。
func (this *LFUCache) incrementFreq(key int) {
	freq := this.keyToFreq[key]
	this.keyToFreq[key] = freq + 1

	// 从当前频率列表中移除该键
	keys := this.freqToKeys[freq]
	for e := keys.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			keys.Remove(e)
			break
		}
	}

	// 如果当前频率列表为空且频率为最小频率，更新最小频率
	if keys.Len() == 0 && freq == this.minFreq {
		this.minFreq++
	}

	// 将该键插入到新的频率列表中
	if this.freqToKeys[freq+1] == nil {
		this.freqToKeys[freq+1] = list.New()
	}
	this.freqToKeys[freq+1].PushBack(key)
}

// 当缓存容量满了时，移除使用频率最小的键（minFreq 对应的键集合中的最久未使用键）。
func (this *LFUCache) removeMinFreqKey() {
	keys := this.freqToKeys[this.minFreq]
	// 移除最小频率列表中的第一个键（最久未使用）
	toRemove := keys.Front().Value.(int)
	keys.Remove(keys.Front())
	if keys.Len() == 0 {
		delete(this.freqToKeys, this.minFreq)
	}
	// 删除该键的记录
	delete(this.keyToValue, toRemove)
	delete(this.keyToFreq, toRemove)
}
