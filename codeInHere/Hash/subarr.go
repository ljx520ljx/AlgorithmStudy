package Hash

//和为k的子数组:给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//子数组是数组中元素的连续非空序列。
/*
算法思路:枚举或前缀和+哈希表优化
外层循环 从 0 遍历到数组的末尾，每个 i 都是子数组的末尾。
内层循环 end 从当前的 start 到 0，向前扫描，计算当前子数组的和。
每次计算子数组的和后，检查它是否等于 k，如果是，就增加计数 count。

pre 表示当前数组从 0 到 i 的前缀和，m 是一个哈希表，用来记录每个前缀和出现的次数。
当前前缀和是 pre，如果存在某个位置 j 使得 pre - pre[j] = k，那么 pre[j] 就是一个有效的前缀和，表示从 j 到 i 的子数组和为 k。
利用哈希表快速查找前缀和出现的次数，从而统计符合条件的子数组。

s 数组存储了所有的前缀和，其中 s[i] 表示数组 nums 中前 i 个元素的和。
cnt 是一个哈希表，用来记录前缀和的出现次数。
遍历 s 数组，对于每个前缀和 sj，如果 sj - k 存在于 cnt 中，说明从某个位置到当前的位置的子数组和为 k，就增加计数 ans。
最后更新哈希表 cnt，记录当前前缀和 sj 的出现次数。
*/

// subarraySum 使用暴力枚举法计算和为 k 的子数组个数
// 时间复杂度: O(N^2)
// 空间复杂度: O(1)
func subarraySum(nums []int, k int) int {
	// count 用于记录和为 k 的子数组的数量
	count := 0

	// 外层循环固定子数组的结束位置 i
	for i := 0; i < len(nums); i++ {
		// sum 用于计算当前子数组的和
		sum := 0
		// 内层循环从结束位置 i 向前遍历，确定子数组的开始位置 end
		// 计算子数组 nums[end...i] 的和
		for end := i; end >= 0; end-- {
			// 累加当前元素到 sum
			sum += nums[end]
			// 如果当前子数组的和等于 k，增加计数
			if sum == k {
				count++
			}
		}
	}
	// 返回最终统计的数量
	return count
}

// subarraySum1 使用前缀和 + 哈希表优化计算和为 k 的子数组个数
// 时间复杂度: O(N)
// 空间复杂度: O(N)
func subarraySum1(nums []int, k int) int {
	// count: 结果计数器
	// pre: 当前遍历到的位置的前缀和 (即 prefixSum[i])
	count, pre := 0, 0
	// m: 哈希表，用于存储每个前缀和出现的次数
	// key: 前缀和的值
	// value: 该前缀和出现的次数
	m := map[int]int{}
	// 初始化：前缀和为 0 (代表空数组或起始位置) 出现了 1 次
	// 这是为了正确处理从数组开头就满足条件的子数组
	m[0] = 1

	// 遍历数组 nums
	for i := 0; i < len(nums); i++ {
		// 更新当前前缀和 pre
		pre += nums[i]
		// 查找是否存在 pre - k 的前缀和
		// 如果 m[pre - k] 存在，说明在 i 之前存在一个或多个 j
		// 使得 prefixSum[j] = pre - k，即 prefixSum[i] - prefixSum[j] = k
		// 子数组 nums[j+1...i] 的和为 k
		if _, ok := m[pre-k]; ok {
			// 将找到的次数累加到 count
			count += m[pre-k]
		}
		// 将当前前缀和 pre 加入哈希表，或增加其出现次数
		m[pre] += 1
	}
	// 返回最终统计的数量
	return count
}

// subarraySum2 使用显式前缀和数组 + 哈希表计算和为 k 的子数组个数
// 时间复杂度: O(N)
// 空间复杂度: O(N)
func subarraySum2(nums []int, k int) (ans int) {
	// s: 前缀和数组，s[i] 表示 nums 前 i 个元素的和 (nums[0...i-1])
	// 长度为 len(nums) + 1，s[0] = 0
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x // 计算前缀和
	}

	// cnt: 哈希表，用于存储遍历过程中遇到的前缀和及其出现次数
	// key: 前缀和的值
	// value: 该前缀和在当前遍历位置之前出现的次数
	cnt := map[int]int{}
	// 遍历前缀和数组 s
	for _, sj := range s { // sj 代表 s[j]
		// 查找是否存在前缀和 si = sj - k (其中 i < j)
		// 如果 cnt[sj - k] > 0，说明在 j 之前存在一个或多个 i
		// 使得 s[i] = sj - k，即 s[j] - s[i] = k
		// 对应子数组 nums[i...j-1] 的和为 k
		ans += cnt[sj-k]
		// 将当前遍历到的前缀和 sj 加入哈希表 cnt，或增加其计数
		// 注意：必须在查找之后再更新 cnt，保证找到的是严格位于 j 之前的前缀和
		cnt[sj]++
	}
	// 返回最终统计的数量 (使用命名返回值 ans)
	return
} 