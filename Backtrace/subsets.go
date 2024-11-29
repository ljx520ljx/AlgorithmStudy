package Backtrace

// 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的 子集 （幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
/*
算法思路:回溯
set 是一个临时的数组，用来保存当前递归路径上的子集。
dfs 是一个递归函数，负责生成子集。在每一步，递归会考虑是否包括当前元素 nums[cur]。
当 cur == len(nums) 时，表示已经处理完所有元素，此时将当前的 set（即子集）添加到结果 ans 中。
递归的核心思想是：
包括当前元素 nums[cur]。
不包括当前元素 nums[cur]。
*/
func subsets(nums []int) (ans [][]int) {
	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

func subsets1(nums []int) [][]int {
	var set []int
	var dfs func(int)
	var res [][]int
	dfs = func(curr int) {
		if curr == len(nums) {
			res = append(res, append([]int(nil), set...))
			return
		}
		set = append(set, nums[curr])
		dfs(curr + 1)
		set = set[:len(set)-1]
		dfs(curr + 1)
	}
	dfs(0)
	return res
}
