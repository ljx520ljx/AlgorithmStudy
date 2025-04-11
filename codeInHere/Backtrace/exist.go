package Backtrace

import (
	"slices"
)

// 全局变量，定义上下左右四个方向的偏移量
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

// 79. 单词搜索
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
/*
算法思路: 深度优先搜索 (DFS) + 回溯 + 优化
1. 统计字符数量（优化一）：
   - 统计 board 中每个字符的出现次数。
   - 统计 word 中每个字符的出现次数。
   - 如果 word 中某个字符的数量超过了 board 中的数量，则可以直接返回 false。
2. 搜索方向优化（优化二）：
   - 比较 word 的首尾字符在 board 中的出现频率。
   - 如果尾字符频率较低，则反转 word 进行搜索。这是一种启发式优化，期望从频率较低的字符开始搜索，可能更快地发现不匹配，从而提前剪枝。
3. DFS + 回溯：
   - 从 board 的每个单元格开始尝试进行 DFS 搜索。
   - 定义 `dfs(i, j, k)` 函数，表示从 board[i][j] 开始，尝试匹配 word[k] 及之后的字符。
   - 标记访问过的单元格：为了防止在当前搜索路径中重复使用同一个单元格，将访问过的单元格临时修改为一个特殊值（这里用 0）。
   - 递归探索：尝试向当前单元格的上下左右四个方向递归调用 dfs。
   - 恢复现场（回溯）：在从当前单元格返回之前，将其值恢复为原来的字符，以便其他搜索路径可以使用它。
4. 边界条件：
   - 坐标越界。
   - 当前 board 单元格字符与 word[k] 不匹配。
   - `k` 到达 word 末尾，表示匹配成功。
*/
func exist(board [][]byte, word string) bool {
	// 统计 board 中每个字符的出现次数
	cnt := map[byte]int{}
	for _, row := range board {
		for _, c := range row {
			cnt[c]++
		}
	}

	// 优化一：检查 board 中的字符数量是否足够构成 word
	w := []byte(word) // 将 word 转为 byte 切片方便操作
	wordCnt := map[byte]int{}
	for _, c := range w {
		wordCnt[c]++
		// 如果 word 中某个字符所需数量 > board 中实际数量，则不可能构成
		if wordCnt[c] > cnt[c] {
			return false
		}
	}

	// 优化二：比较首尾字符频率，尝试从频率低的字符开始搜索
	if cnt[w[len(w)-1]] < cnt[w[0]] {
		slices.Reverse(w) // 反转 word 切片
	}

	m, n := len(board), len(board[0]) // 获取 board 的尺寸
	var dfs func(int, int, int) bool  // 声明 dfs 函数变量

	// 定义 DFS 函数：(当前行 i, 当前列 j, 当前匹配到 word 的索引 k)
	dfs = func(i, j, k int) bool {
		// 1. 越界或字符不匹配，此路不通
		// 注意：这里 board[i][j] != w[k] 也处理了 board[i][j] 被标记为 0 的情况
		if board[i][j] != w[k] {
			return false
		}
		// 2. 匹配到 word 的最后一个字符，成功
		if k == len(w)-1 {
			return true
		}

		// 3. 标记当前单元格已在本次搜索路径中访问过
		//    将其值修改为一个不可能出现在 word 中的值（例如 0）
		originalChar := board[i][j] // 先保存原始字符，用于恢复
		board[i][j] = 0

		// 4. 探索相邻单元格
		for _, d := range dirs {
			x, y := i+d.x, j+d.y // 计算相邻单元格坐标
			// 检查相邻坐标是否在 board 范围内，并递归搜索
			if 0 <= x && x < m && 0 <= y && y < n && dfs(x, y, k+1) {
				// 如果从某个相邻单元格出发找到了完整路径，则返回 true
				// 在返回 true 前需要恢复现场吗？
				// 不需要，因为一旦找到，我们就不关心 board 的状态了，直接层层返回 true 即可。
				// 但是为了代码逻辑一致性和可维护性，或者如果需要在找到后继续找其他可能（虽然本题不需要），恢复现场总是好习惯。
				// 严格来说，这里可以直接 return true，但为了下面恢复现场逻辑的完整性，先不直接返回。
				// 更新：考虑到找到就结束，这里的恢复不是必须的。但为了逻辑清晰，标准回溯会恢复。
				// 为了保持下面恢复现场代码的统一性，这里也先不直接return
				// ---- 标准回溯代码应该在这里恢复现场并返回true ----
				// board[i][j] = originalChar // (可选，如果找到就结束则非必须)
				return true // 搜到了！
			}
		}

		// 5. 恢复现场（回溯）
		// 如果四个方向都没有找到路径，说明从 (i, j) 开始匹配 word[k] 的尝试失败
		// 需要将 board[i][j] 恢复为原始字符，以便其他搜索路径可以使用它
		board[i][j] = originalChar
		// 6. 返回 false，表示从当前路径没搜到
		return false
	}

	// 从 board 的每个单元格开始尝试 DFS
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) { // 从 (i, j) 开始匹配 word[0]
				return true // 只要找到一个起点能完成搜索，就返回 true
			}
		}
	}

	// 如果所有起点都尝试过，仍未找到，则返回 false
	return false
}
func exist1(board [][]byte, word string) bool {
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
