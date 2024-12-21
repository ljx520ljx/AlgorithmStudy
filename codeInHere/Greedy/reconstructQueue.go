package Greedy

//406. 根据身高重建队列
/*
算法思路:贪心算法
首先对输入的 people 数组进行特殊排序，保证身高高的人在前面，身高相同的情况下，k 值小的人在前面。
这样在插入元素时，我们可以根据 k 值将元素插入到结果数组的相应位置，
因为前面已经插入的元素都是比当前元素高或者相等的，且插入位置是按照 k 值计算的，能够保证插入后满足条件。
对于每个元素，将其插入到结果数组中 k 值对应的位置，同时将该位置之后的元素后移一位，最终得到满足条件的队列。
k表示前面有多少人比他高或者相等
时间复杂度：O(n^2)，其中 n 是 people 数组的长度。排序的时间复杂度为 O(nlogn)，插入元素的时间复杂度为 O(n^2)。
空间复杂度：O(nlogn)，即为排序需要使用的栈空间。
*/

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// 首先按照身高降序排序，如果身高相同则按照前面有更高或相同身高的人的数量升序排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	// 结果切片
	result := make([][]int, len(people))
	// 遍历排序后的 people 切片
	for _, person := range people {
		index := person[1]
		// 将当前元素插入到结果切片的指定位置，将后面元素后移
		copy(result[index+1:], result[index:])
		result[index] = person
	}
	return result
}
