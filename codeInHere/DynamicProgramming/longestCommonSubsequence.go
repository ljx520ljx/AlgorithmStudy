package DynamicProgramming

//1143. 最长公共子序列
//给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
/*
算法思路:动态规划(空间优化版)
1.确保text1是较长的字符串，这样可以使用较短的字符串长度作为dp数组长度
2.创建一维dp数组，长度为较短字符串长度+1
3.使用pre变量记录左上角的值(对应二维dp中的dp[i-1][j-1])
4.使用temp变量暂存当前值，供下一轮迭代使用
5.状态转移：
  - 当前字符相同时：dp[j] = pre + 1
  - 当前字符不同时：dp[j] = max(dp[j], dp[j-1])
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	// 确保text1是较长的字符串，优化空间使用
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}
	m, n := len(text1), len(text2)
	// 创建一维dp数组，使用较短字符串的长度
	dp := make([]int, n+1)

	// 遍历较长的字符串
	for i := 1; i <= m; i++ {
		// pre记录左上角的值（对应二维dp中的dp[i-1][j-1]）
		pre := 0
		// 遍历较短的字符串
		for j := 1; j <= n; j++ {
			// 暂存当前值，供下一轮迭代使用
			temp := dp[j]
			// 当前字符相同时
			if text1[i-1] == text2[j-1] {
				// 使用左上角的值+1
				dp[j] = pre + 1
			} else {
				// 当前字符不同时，取左边或上边的较大值
				dp[j] = max(dp[j], dp[j-1])
			}
			// 更新左上角的值
			pre = temp
		}
	}
	return dp[n]
}

// 二维dp在此:
func longestCommonSubsequence1(text1 string, text2 string) int {
	// 1. 获取两个字符串的长度
	m, n := len(text1), len(text2)

	// 2. 创建dp数组，dp[i][j]表示text1[0:i]和text2[0:j]的最长公共子序列长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 3. 动态规划过程
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果当前字符相同
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 如果当前字符不同，取较大值
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
