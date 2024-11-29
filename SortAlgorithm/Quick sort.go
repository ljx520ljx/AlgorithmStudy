package SortAlgorithm

import (
	"math/rand"
)

//手撕快速排序
/*
算法思路:
选择基准元素：从数组中选择一个元素作为基准元素，使用随机种子。
分区操作：三向切分,重新排列数组，使得基准元素左边的所有元素都比它小，
右边的所有元素都比它大。这个过程被称为“分区”操作。
递归排序：对基准元素左边和右边的子数组递归进行相同的操作。
终止条件：当子数组的长度为 1 或 0 时，递归结束。
*/

// 快速排序函数
// 快速排序的主函数
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 随机基准选择 + 三向切分的快速排序递归函数
func quickSort(nums []int, low, high int) {
	if low < high {
		// 随机选择基准并交换到low位置
		randIndex := rand.Intn(high-low+1) + low
		nums[low], nums[randIndex] = nums[randIndex], nums[low]

		// 三向切分分割数组
		lt, gt := threeWayPartition(nums, low, high)

		// 对小于基准部分递归排序
		quickSort(nums, low, lt-1)
		// 对大于基准部分递归排序
		quickSort(nums, gt+1, high)
	}
}

// 三向切分函数
func threeWayPartition(nums []int, low, high int) (int, int) {
	pivot := nums[low]
	lt := low  // 小于基准的区域的右边界
	gt := high // 大于基准的区域的左边界
	i := low + 1

	for i <= gt {
		if nums[i] < pivot {
			// 当前元素小于基准，放到左边
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++
		} else if nums[i] > pivot {
			// 当前元素大于基准，放到右边
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			// 当前元素等于基准，跳过
			i++
		}
	}

	// 返回三部分的边界
	return lt, gt
}
