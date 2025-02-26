package SortAlgorithm

//274. H 指数
//给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。
/*
算法思路:计数排序
通过统计每个引用次数出现的次数，将问题转化为对引用次数的分布进行分析。
从高到低遍历引用次数，累加满足条件的论文数量，直到找到最大的 h。
从高到低遍历引用次数，而不是从低到高。用空间减少时间复杂度。
*/
func hIndex(citations []int) int {
	n := len(citations)
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		citations[i] = min(citations[i], n)
		cnt[citations[i]]++
	}
	s := 0
	for i := n; i >= 0; i-- {
		s += cnt[i]
		if s >= i {
			return i
		}
	}
	return 0
}
