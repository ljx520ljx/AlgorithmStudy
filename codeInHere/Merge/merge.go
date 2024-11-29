package Merge

//合并区间以数组 intervals 表示若干个区间的集合，其中单个区间为intervals[i]=[starti, endi]
//请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间.
/*
算法思路:先排序,再合并
先把每个区间按左端点大小从小到大排序
接着对比右端点进行合并
*/

import (
	"slices"
	"sort"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })
	var res [][]int
	for _, v := range intervals {
		n := len(res)
		if n > 0 && v[0] <= res[n-1][1] {
			res[n-1][1] = max(res[n-1][1], v[1])
		} else {
			res = append(res, v)
		}
	}
	return res
}

func merge1(intervals [][]int) [][]int {
	// 按照区间的起始值进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}

	for _, interval := range intervals {
		L, R := interval[0], interval[1]

		// 如果 merged 是空的，或者当前区间与最后一个区间不重叠
		if len(merged) == 0 || merged[len(merged)-1][1] < L {
			merged = append(merged, []int{L, R})
		} else {
			// 合并重叠的区间
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], R)
		}
	}

	return merged
}
