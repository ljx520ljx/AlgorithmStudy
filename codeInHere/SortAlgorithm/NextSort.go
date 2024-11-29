package SortAlgorithm

//下一个排序
/*
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。
更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，
那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。
如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
*/
/*
算法思路:
1.从右往左找第一个递增对(nums[i], nums[i+1]),这样剩下的nums[i+1:] 部分是降序排列的。
2.当找到了这样的i,再从数组末尾开始，依次找到第一个大于 nums[i] 的元素 nums[j]。
3.交换nums[i] 和 nums[j]： 交换这两个元素，以得到比当前排列更大的数。
4.反转 i 之后的子数组：因为剩下的nums[i+1:]是降序排列,反转得到的升序排列是最小的。
*/

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// Step 1: 从后往前找到第一个递增对（nums[i] < nums[i+1]）
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 如果找到了这样的 i
	if i >= 0 {
		// Step 2: 从后往前找到第一个大于 nums[i] 的元素 nums[j]
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		// Step 3: 交换 nums[i] 和 nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 4: 反转 i 后面的所有元素
	reverse(nums, i+1, n-1)
}

// 反转 nums[l] 到 nums[r] 区间的元素
func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
