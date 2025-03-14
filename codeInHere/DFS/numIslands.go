package DFS

//岛屿数量
/*
算法思路:深度优先遍历
假如碰到第一片陆地,count++,在把这片陆地连接的所有土地记为水.
这样当下次碰到的新陆地就属于另一片岛屿,count++
直到遍历完整个二维网格
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	var dfs func(i, j int)
	m := len(grid)
	n := len(grid[0])
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

//广度优先搜索
/*
算法思路:
*/
func numIslands1(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	//directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 四个方向
	count := 0

	// 广度优先搜索
	bfs := func(x, y int) {
		queue := [][2]int{{x, y}}
		for len(queue) > 0 {
			// 出队
			curr := queue[0]
			queue = queue[1:]

			cx, cy := curr[0], curr[1]

			// 上
			if cx-1 >= 0 && grid[cx-1][cy] == '1' {
				grid[cx-1][cy] = '0'
				queue = append(queue, [2]int{cx - 1, cy})
			}
			// 下
			if cx+1 < m && grid[cx+1][cy] == '1' {
				grid[cx+1][cy] = '0'
				queue = append(queue, [2]int{cx + 1, cy})
			}
			// 左
			if cy-1 >= 0 && grid[cx][cy-1] == '1' {
				grid[cx][cy-1] = '0'
				queue = append(queue, [2]int{cx, cy - 1})
			}
			// 右
			if cy+1 < n && grid[cx][cy+1] == '1' {
				grid[cx][cy+1] = '0'
				queue = append(queue, [2]int{cx, cy + 1})
			}
		}
	}

	// 遍历网格
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				// grid[i][j] = '0' // 标记为已访问
				bfs(i, j) // 启动 BFS
			}
		}
	}

	return count
}
