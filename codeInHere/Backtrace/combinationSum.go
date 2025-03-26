package Backtrace

import "slices"

//39. 组合总和
//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。
//candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
/*
算法思路:回溯算法（DFS）
1.对数组排序，便于剪枝和保持结果有序
2.使用path数组记录当前路径，res数组记录所有合法组合
3.使用start参数控制搜索起点，避免重复组合
4.状态转移：
  - 当和大于target时，终止当前路径
  - 当和等于target时，保存当前组合
  - 从start开始尝试每个数字，可以重复使用当前数字
5.回溯时移除最后添加的数字，恢复状态
*/

func combinationSum(candidates []int, target int) (res [][]int) {
	slices.Sort(candidates)
	n := len(candidates)
	var set []int
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if target == 0 {
			res = append(res, slices.Clone(set))
			return
		}

		if index >= n || target < candidates[index] {
			return
		}
		//不选跳过
		dfs(target, index+1)
		//选
		set = append(set, candidates[index])
		dfs(target-candidates[index], index)
		set = set[:len(set)-1]

	}

	dfs(target, 0)
	return res
}

func combinationSum1(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	var path []int
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			// 找到一个合法组合
			ans = append(ans, slices.Clone(path))
			return
		}

		if left < candidates[i] {
			return
		}

		// 枚举选哪个
		for j := i; j < len(candidates); j++ {
			path = append(path, candidates[j])
			dfs(j, left-candidates[j])
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0, target)
	return ans
}

func combinationSum2(candidates []int, target int) (res [][]int) {
	n := len(candidates)
	var set []int
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index >= n {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), set...))
			return
		}
		//不选跳过
		dfs(target, index+1)
		//选
		if target-candidates[index] >= 0 {
			set = append(set, candidates[index])
			dfs(target-candidates[index], index)
			set = set[:len(set)-1]
		}
	}

	dfs(target, 0)
	return
}

func combinationSum3(candidates []int, target int) [][]int {
	// 排序数组，使结果有序
	slices.Sort(candidates)
	// 结果数组和路径数组
	var res [][]int
	var path []int

	// 定义DFS函数，start表示搜索起点，sum表示当前和
	var dfs func(start, sum int)
	dfs = func(start, sum int) {
		// 和超过目标值，剪枝
		if sum > target {
			return
		}
		// 找到合法组合，保存结果
		if sum == target {
			res = append(res, slices.Clone(path))
			return
		}

		// 从start开始尝试每个数字
		for i := start; i < len(candidates); i++ {
			// 将当前数字加入路径
			path = append(path, candidates[i])
			// 继续DFS，可以重复使用当前数字，所以传入i而不是i+1
			dfs(i, sum+candidates[i])
			// 回溯，移除最后添加的数字
			path = path[:len(path)-1]
		}
	}

	// 从位置0开始搜索，初始和为0
	dfs(0, 0)
	return res
}
