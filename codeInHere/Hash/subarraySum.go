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

func subarraySum(nums []int, k int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		//以i为结尾的每一种组合
		for end := i; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySum1(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func subarraySum2(nums []int, k int) (ans int) {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	cnt := map[int]int{}
	for _, sj := range s {
		ans += cnt[sj-k]
		cnt[sj]++
	}
	return
}
