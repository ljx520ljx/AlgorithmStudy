package DynamicProgramming

//编辑距离
//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数
//你可以对一个单词进行如下三种操作：插入一个字符 删除一个字符 替换一个字符
/*
算法思路:动态规划
定义 f[i] 为将 word1 的前 i 个字符转换成 word2 的前 j 个字符所需的最少操作数。
设 f[i][j] 表示将 word1 的前 i 个字符转换成 word2 的前 j 个字符所需的最小操作数。我们要求的是 f[len(word1)][len(word2)]。
状态转移：
如果 word1[i-1] 和 word2[j-1] 相同，则无需操作，即 f[i][j] = f[i-1][j-1]。
否则，我们需要执行以下三种操作之一：
插入一个字符到 word1 中，继续 word1 的前 i 个字符与 word2 的前 j-1 个字符比较，即 f[i][j] = f[i][j-1] + 1。
删除一个字符，从 word1 中删除字符，继续 word1 的前 i-1 个字符与 word2 的前 j 个字符的比较，即 f[i][j] = f[i-1][j] + 1。
替换一个字符，替换 word1[i-1] 为 word2[j-1]，变成 f[i-1][j-1] + 1。
初始化：
由于可以对单词进行删除、插入操作，第一行和第一列需要初始化为逐步删除和插入字符的最少操作数：
f[0][j] = j，表示将空字符串转换成 word2 的前 j 个字符需要 j 次插入操作。
f[i][0] = i，表示将 word1 的前 i 个字符转换为空字符串需要 i 次删除操作。
时间复杂度：O(nm)，其中 n 为 s 的长度，m 为 t 的长度。
*/

func minDistance3(s, t string) int {
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

func minDistance(s, t string) int {
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j <= m; j++ {
		f[0][j] = j
	}
	for i, x := range s {
		f[i+1][0] = i + 1
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}

// 使用滚动数组优化，只维护两行状态：f[i%2] 和 f[(i+1)%2]。
func minDistance1(s, t string) int {
	n, m := len(s), len(t)
	f := [2][]int{make([]int, m+1), make([]int, m+1)}
	for j := 1; j <= m; j++ {
		f[0][j] = j
	}
	for i, x := range s {
		f[(i+1)%2][0] = i + 1
		for j, y := range t {
			if x == y {
				f[(i+1)%2][j+1] = f[i%2][j]
			} else {
				f[(i+1)%2][j+1] = min(f[i%2][j+1], f[(i+1)%2][j], f[i%2][j]) + 1
			}
		}
	}
	return f[n%2][m]
}

// 利用一维数组 f 来存储当前行的编辑距离，并且通过 pre 变量保存上一行的状态来减少空间复杂度。
func minDistance2(s, t string) int {
	m := len(t)
	f := make([]int, m+1)
	for j := 1; j <= m; j++ {
		f[j] = j
	}
	for _, x := range s {
		pre := f[0]
		f[0]++ // f[0] = i + 1
		for j, y := range t {
			if x == y {
				f[j+1], pre = pre, f[j+1]
			} else {
				f[j+1], pre = min(f[j+1], f[j], pre)+1, f[j+1]
			}
		}
	}
	return f[m]
}
