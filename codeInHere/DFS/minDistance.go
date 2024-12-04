package DFS

//编辑距离
//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数
//你可以对一个单词进行如下三种操作：插入一个字符 删除一个字符 替换一个字符
/*
算法思路:递归搜索 + 保存计算结果 = 记忆化搜索
dp[i][j]，表示将 s[0:i] 转换为 t[0:j] 所需的最少操作数。
基本情况：如果 i < 0，返回 j + 1（插入操作）；如果 j < 0，返回 i + 1（删除操作）。
通过 memo[i][j] 来存储已经计算过的结果
如果字符 s[i] == t[j]，跳过该字符，递归调用 dfs(i-1, j-1)。
否则，计算三种操作（删除、插入、替换）的最小值并加 1。
时间复杂度：O(n * m)，其中 n 和 m 是字符串 s 和 t 的长度。
空间复杂度：O(n * m)，用于存储 memo 数组。
*/

func minDistance(s, t string) int {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示还没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if s[i] == t[j] {
			return dfs(i-1, j-1)
		}
		return min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
	}
	return dfs(n-1, m-1)
}
