package DynamicProgramming

/*组合总数
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。任意顺序返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
*/
/*
算法思路:动态规划
*/

func combinationSum(candidates []int, target int) [][]int {
	dp := make([][][]int, target+1)
	dp[0] = [][]int{{}}
	for i := 0; i < len(candidates); i++ {
		for j := candidates[i]; j < len(dp); j++ {
			for _, group := range dp[j-candidates[i]] {
				newGroup := append([]int{}, group...)
				newGroup = append(newGroup, candidates[i])
				dp[j] = append(dp[j], newGroup)
			}
		}
	}
	return dp[target]
}
