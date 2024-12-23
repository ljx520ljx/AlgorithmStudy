package MinHeap

//347. 前 K 个高频元素
/*
算法思路:最小堆
频率统计：创建一个 map[int]int 类型的 occurrences 映射，用于存储输入数组 nums 中每个元素及其出现的频率。
最小堆：创建一个最小堆 h，用于存储元素及其出现频率。堆的每个元素是一个 [2]int 数组，其中 [0] 表示元素的值，[1] 表示元素的出现频率。
遍历映射：遍历 occurrences 映射，对于每个元素及其出现频率，将其添加到最小堆 h 中。如果堆的大小超过了 k，就移除堆顶元素，以保证堆中始终保留前 k 个高频元素。
结果获取：堆中剩余的 k 个元素即为前 k 个高频元素。将堆中的元素从堆顶依次取出，并将其值添加到结果数组中。
*/
import (
	"container/heap"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
算法思路2:hash表+频率映射存储
频率统计:使用一个 map[int]int 类型的 numMap 来统计输入数组 nums 中每个元素的出现频率。
频率映射存储：创建一个二维切片 valueMap，用于存储每个元素的出现频率及其对应的元素值。
有效频率筛选：初始化 freqs 切片，用于存储有效频率。遍历 valueMap，将不为 nil 的元素对应的频率添加到 freqs 中.
频率排序：对 freqs 进行降序排序。
元素选取：遍历 freqs 切片，对于每个频率，从 valueMap 中取出相应元素添加到 res 切片中。当添加元素的数量达到 k 时，停止添加并返回结果。
*/
func topKFrequent1(nums []int, k int) (res []int) {
	numMap := make(map[int]int, len(nums))
	valueMap := make([][]int, len(nums)+1)
	for _, num := range nums {
		numMap[num]++
	}
	for num, freq := range numMap {
		valueMap[freq] = append(valueMap[freq], num)
	}
	var freqs []int
	for freq, numsSlices := range valueMap {
		if numsSlices != nil {
			freqs = append(freqs, freq)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))
	for _, freq := range freqs {
		for _, num := range valueMap[freq] {
			if k > 0 {
				res = append(res, num)
				k--
			} else {
				return
			}
		}
	}
	return
}
