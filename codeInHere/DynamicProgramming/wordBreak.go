package DynamicProgramming

//单词拆分
//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
/*
算法思路:动态规划
定义动态规划状态 dp[i]：
dp[i] 表示字符串 s 的前 i 个字符是否可以拆分成字典中存在的若干个单词。
初始化时，dp[0] = true，因为空字符串可以视为合法的拆分（没有任何字符，当然合法）。
状态转移：
对于每个 i（表示当前考虑到的字符串长度），我们检查 dp[i] 的值，是否可以由之前的某个位置 j 递推过来。
递推的条件是，s[j..i-1] 是否是字典中的一个单词，同时 dp[j] 需要为 true，表示 s[0..j-1] 这个部分已经合法地拆分成字典中的单词。
转移方程：
dp[i]=dp[j]&&wordSet[s[j:i]]
其中 check(s[j..i-1]) 判断子串 s[j..i-1] 是否存在于字典中。
剪枝：
对于 s[j..i-1] 的检查，如果 i - j 的长度已经大于字典中任何一个单词的长度，那么就没有必要继续枚举 j，可以提前结束。
*/

func wordBreak(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	f := make([]bool, n+1) //n+1是因为考虑空字符串所以要+1
	f[0] = true            //表示空字符串
	for i := 1; i <= n; i++ {
		//这里j >= max(i-maxLen, 0)用了剪枝优化
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if f[j] && words[s[j:i]] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}
