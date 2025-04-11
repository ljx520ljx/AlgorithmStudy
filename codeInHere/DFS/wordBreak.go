package DFS

// 139. 单词拆分
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
/*
算法思路: 动态规划 (DP) (自底向上)
1. 状态定义：
   - dp[i] 表示字符串 s 的前 i 个字符组成的子串 s[0:i] 是否可以被字典中的单词成功拆分。
2. 基础情况：
   - dp[0] = true （空字符串可以被拆分）。
3. 状态转移：
   - 遍历 i 从 1 到 n (字符串长度)。
   - 对于每个 i，遍历 j 从 0 到 i-1。
   - 如果 dp[j] 为 true (表示 s[0:j] 可拆分) 并且子串 s[j:i] 在字典中存在，
     则 dp[i] = true，并可以 break 内层循环（因为只要找到一种拆分方式即可）。
4. 最终结果：
   - 返回 dp[n]。
5. 预处理：
   - 将 wordDict 转换为哈希集合，方便快速查找。
*/
func wordBreak(s string, wordDict []string) bool {
	// 预处理：将字典转换为哈希集合，提高查找效率 O(L), L为字典总长度
	wordDictSet := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	n := len(s)
	// 创建 DP 数组，长度为 n+1
	dp := make([]bool, n+1)
	// 基础情况：空字符串总是可以被拆分
	dp[0] = true

	// 动态规划主循环：计算 dp[1] 到 dp[n]
	// i 表示当前考虑的子串 s[0:i] 的长度
	for i := 1; i <= n; i++ {
		// 遍历所有可能的分割点 j
		// j 表示 s[0:j] 的长度，而 s[j:i] 是尝试匹配的最后一个单词
		for j := 0; j < i; j++ {
			// 获取可能的最后一个单词子串
			substring := s[j:i]
			// 状态转移条件：
			// 1. s[0:j] 必须能够被拆分 (dp[j] == true)
			// 2. 子串 s[j:i] 必须是字典中的单词 (wordDictSet[substring])
			if dp[j] && wordDictSet[substring] {
				// 如果满足条件，则 s[0:i] 可以被拆分
				dp[i] = true
				// 优化：一旦找到一种拆分方式，就不需要再检查其他 j 了
				break
			}
		}
		// 如果内层循环结束，dp[i] 仍然是 false，说明找不到任何 j 能满足条件
	}

	// 返回整个字符串 s (即 s[0:n]) 的拆分结果
	return dp[n]
}
