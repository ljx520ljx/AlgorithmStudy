package DynamicProgramming

/*
正则表达式匹配
给你一个字符串s和一个字符规律p,请你来实现一个支持'.'和'*'的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s 的，而不是部分字符串。
算法思路:二维动态规划
使用二维表 dp[i][j]表示字符串s的前i个字符与模式p的前 j 个字符是否匹配。
状态转移方程
如果 p[j-1] 是普通字符或 .：
dp[i][j] = dp[i-1][j-1]，且 s[i-1] == p[j-1] 或 p[j-1] == '.'
如果 p[j-1] 是 *：
可以匹配零个前一字符：dp[i][j]=dp[i][j-2]
或匹配一个或多个前一字符:dp[i][j]=dp[i-1][j]且s[i-1]==p[j-2]或p[j-2]=='.'
边界条件
dp[0][0] = true：空字符串和空模式匹配。
当字符串为空时,模式中每个*可能匹配前一个字符零次,因此需要初始化对应的dp。
时间复杂度
时间复杂度：O(m * n)，m 为字符串 s 的长度，n 为模式串 p 的长度。
空间复杂度：O(m * n)。
*/

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 空字符串和空模式匹配
	dp[0][0] = true

	// 初始化当字符串为空时的匹配情况
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// 动态规划
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				//下面这句话可读性太差,所以改了一下
				//dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (p[j-2] == s[i-1] || p[j-2] == '.'))
				// 匹配零个前一字符
				zeroMatch := dp[i][j-2]

				// 匹配一个或多个前一字符
				oneOrMoreMatch := dp[i-1][j] && (p[j-2] == s[i-1] || p[j-2] == '.')

				// 综合两种可能
				dp[i][j] = zeroMatch || oneOrMoreMatch
			}
		}
	}

	return dp[m][n]
}
