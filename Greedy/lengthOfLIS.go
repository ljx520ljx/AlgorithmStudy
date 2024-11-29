package Greedy

//最长递增子序列
/*
算法思路:贪心+二分查找
核心思想：
使用一个数组 d，其中 d[i] 表示长度为 i+1i+1i+1 的递增子序列的最小可能末尾元素。
d 的长度即为当前找到的最长递增子序列的长度。
贪心策略：
每次遇到一个新数 num，希望它尽可能小地插入到递增子序列中，以便为后续的数字留出更大的空间。
二分查找：
对于每个 num，用二分查找确定它应该插入到 d 中的位置。
如果 num 比 d 中的所有元素都大，就扩展 d。
否则，用 num 替换 d 中第一个大于等于 num 的元素
(这样得到的递增子序列会改变,但是长度不变,相当于用来辅助计算长度)保证 d 的单调递增性。
时间复杂度：
遍历数组的时间复杂度为O(n)。
对每个元素进行二分查找的时间复杂度为O(logn)。
总时间复杂度为O(nlogn)。
空间复杂度：
使用了一个数组 d 来存储递增子序列的最小末尾元素，空间复杂度为O(n)。
*/
func lengthOfLIS(nums []int) int {
	// 定义一个数组 d 用于存储当前递增子序列的最小末尾元素
	d := []int{}

	for _, num := range nums {
		// 使用二分查找找到 d 中第一个大于等于 num 的位置
		left, right := 0, len(d)
		for left < right {
			mid := (left + right) / 2
			if d[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 如果 num 比 d 中所有的元素都大，则扩展 d
		if left == len(d) {
			d = append(d, num)
		} else {
			// 否则更新 d[left] 的值
			d[left] = num
		}
	}

	// d 的长度即为最长递增子序列的长度
	return len(d)
}
