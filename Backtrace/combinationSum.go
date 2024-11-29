package Backtrace

import "slices"

/*组合总数
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。任意顺序返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
*/
/*
算法思路:回溯
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
	n:=len(candidates)
	var set []int
	var dfs func(target,index int)
	dfs=func(target,index int){
		if index>=n{
			return
		}
		if target==0{
			res=append(res,append([]int(nil),set...))
			return
		}
		//不选跳过
		dfs(target,index+1)
		//选
		if target-candidates[index] >= 0{
			set=append(set,candidates[index])
			dfs(target-candidates[index],index)
			set=set[:len(set)-1]
		}
	}

	dfs(target,0)
	return
}