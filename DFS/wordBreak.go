package DFS

//单词拆分
//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
/*
算法思路:dfs
*/

func wordBreak(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	memo := make([]int8, n+1)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int8
	dfs = func(i int) (res int8) {
		if i == 0 { // 成功拆分！
			return 1
		}
		p := &memo[i]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if words[s[j:i]] && dfs(j) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(n) == 1
}
