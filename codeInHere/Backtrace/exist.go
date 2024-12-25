package Backtrace

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
