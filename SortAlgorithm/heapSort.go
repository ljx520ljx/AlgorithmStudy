package SortAlgorithm

// 手撕堆排序
/*
算法思路:
构建最大堆:
将数组转换成一个最大堆，从数组的中间元素（第一个非叶节点）开始向上调整堆的结构。
排序过程:
每次将堆顶元素（最大值）与最后一个元素交换，并缩小堆的有效范围。
对交换后的堆重新调整，使其满足堆的性质。
重复上述操作，直到所有元素排序完成。
heapSort 函数:
构建最大堆的时间复杂度是O(n)。
排序阶段每次调整堆的时间复杂度是O(logn)，总共需要调整n−1 次，因此排序阶段的时间复杂度是O(nlogn)。
heapify 函数:
维护最大堆性质的递归函数，通过比较父节点和子节点的大小，将最大值调整到堆顶。
交换和调整:
每次交换堆顶和最后一个元素后，缩小堆范围并继续调整，最终实现排序。
*/

func sortArray1(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	n := len(nums)

	// 1. 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	// 2. 逐步提取堆顶元素并调整堆
	for i := n - 1; i > 0; i-- {
		// 将堆顶元素与当前尾部元素交换
		nums[0], nums[i] = nums[i], nums[0]
		// 缩小堆的范围，并调整堆
		heapify(nums, i, 0)
	}
}

// 调整堆的函数（维护最大堆性质）
func heapify(nums []int, n, i int) {
	largest := i     // 当前节点
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	// 如果左子节点大于当前节点，更新最大值索引
	if left < n && nums[left] > nums[largest] {
		largest = left
	}

	// 如果右子节点大于当前最大值，更新最大值索引
	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	// 如果最大值不是当前节点，交换并继续调整
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		// 递归调整受影响的子树
		heapify(nums, n, largest)
	}
}
