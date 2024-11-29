package DynamicProgramming

// 最长有效括号
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
/*
算法思路:动态规划
状态转移
仅当s[i]为)时 dp[i]才可能有值，需状态转移
分类讨论
当s[i−1]为'('时
dp[i]=dp[i−2]+2

当s[i−1]为')'时
设j=i−1−dp[i−1] ，表示跨越左侧相邻有效括号整体后的第一个下标
当 j>=0,s[j]='(' 时，dp[i]=dp[j−1]+dp[i−1]+2 ，表示跨越左侧相邻整体后左侧的有效长度+左侧整体长度+一对括号
*/
func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return 0
	}

	dp := make([]int, n)
	maxLength := 0

	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { // 处理 "()" 的情况
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] += dp[i-2] // 如果前面有有效括号，累加其长度
				}
			} else if j := i - dp[i-1] - 1; j >= 0 && s[j] == '(' { // 处理 "..." + "()" 的情况
				dp[i] = dp[i-1] + 2
				if j-1 >= 0 {
					dp[i] += dp[j-1] // 累加前面部分的有效括号长度
				}
			}
			maxLength = max(maxLength, dp[i])
		}
	}

	return maxLength
}

// 简写版
// dp
func longestValidParentheses1(s string) int {
	var (
		n   = len(s)
		dp  = make([]int, n+1) // 以s[i]结尾的最长子串长度  整体右移1位兼容边界
		ans int
	)
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}

		// )
		//分类讨论
		if s[i-1] == '(' {
			dp[i+1] = dp[i-1] + 2
		} else {
			j := i - 1 - dp[i] // 跨越左侧相邻整体后的第一个下标
			if j >= 0 && s[j] == '(' {
				dp[i+1] = dp[j] + dp[i] + 2
			}
		}
		ans = max(ans, dp[i+1])
	}

	return ans
}
