package DynamicProgramming

//单词拆分
//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
/*
算法思路:动态规划
自底向上计算。
具体来说，f[i] 的定义和 dfs(i) 的定义是一样的，都表示能否把前缀 s[:i]（表示 s[0] 到 s[i−1]）划分成若干段，使得每段都在 wordDict 中。
同样地，枚举 j=i−1,i−2,i−3,…,max(i−maxLen,0)，只要其中一个 j 满足 s[j:i] 在 wordDict 中且 f[j]=true，那么 f[i] 就是 true。
初始值 f[0]=true，翻译自递归边界 dfs(0)=true。
答案为 f[n]，翻译自递归入口 dfs(n)。
*/

func wordBreak(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if f[j] && words[s[j:i]] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}

func wordBreak1(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
