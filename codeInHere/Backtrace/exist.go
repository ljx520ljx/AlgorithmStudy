package Backtrace

//79. 单词搜索
/*
算法思路:回溯
创建一个与 board 相同大小的二维布尔型矩阵 visited 用于标记已访问的单元格，初始化为 false。
完成匹配检查：检查 index 是否等于 word 的长度，若相等则返回 true。
条件判断与终止：检查边界、是否已访问及字符是否匹配，若任一条件不满足则返回 false。
标记与递归搜索：标记当前单元格已访问，对相邻单元格递归调用 dfs 并将结果存储在 found 中。
回溯操作：将当前单元格标记为未访问以进行回溯。
结果返回：返回 found 表示从当前单元格能否找到匹配路径。
深度优先搜索：从每个起始位置调用 dfs 函数，使用深度优先搜索算法开始尝试匹配 word,找到匹配，立即返回 true。
*/
import "slices"

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(int, int, int) bool
	dfs = func(row int, col int, index int) bool {
		// 检查是否已经匹配完整个单词
		if index == len(word) {
			return true
		}
		// 检查边界条件
		if row < 0 || row >= m || col < 0 || col >= n || visited[row][col] || board[row][col] != word[index] {
			return false
		}
		visited[row][col] = true
		// 递归搜索相邻单元格
		found := dfs(row-1, col, index+1) || dfs(row+1, col, index+1) || dfs(row, col-1, index+1) || dfs(row, col+1, index+1)
		visited[row][col] = false
		return found
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func exist1(board [][]byte, word string) bool {
	cnt := map[byte]int{}
	for _, row := range board {
		for _, c := range row {
			cnt[c]++
		}
	}

	// 优化一
	w := []byte(word)
	wordCnt := map[byte]int{}
	for _, c := range w {
		wordCnt[c]++
		if wordCnt[c] > cnt[c] {
			return false
		}
	}

	// 优化二
	if cnt[w[len(w)-1]] < cnt[w[0]] {
		slices.Reverse(w)
	}

	m, n := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != w[k] { // 匹配失败
			return false
		}
		if k == len(w)-1 { // 匹配成功
			return true
		}
		board[i][j] = 0 // 标记访问过
		for _, d := range dirs {
			x, y := i+d.x, j+d.y // 相邻格子
			if 0 <= x && x < m && 0 <= y && y < n && dfs(x, y, k+1) {
				return true // 搜到了！
			}
		}
		board[i][j] = w[k] // 恢复现场
		return false       // 没搜到
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true // 搜到了！
			}
		}
	}
	return false // 没搜到
}
