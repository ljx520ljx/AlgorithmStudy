package BFS

//岛屿的数量
/*
算法思路:广度优先遍历
遍历网格：
遍历网格中的每个元素。
如果遇到陆地 ('1')，计数器加 1，并将该陆地添加到队列中，开始广度优先搜索。
BFS 实现：
从队列中取出当前节点，并访问其上下左右的相邻节点。
如果相邻节点是陆地 ('1')，将其标记为已访问（改为 '0'），并将其加入队列。
标记已访问：
访问过的陆地 ('1') 改为水 ('0')，防止重复计算。
岛屿计数：
每次遇到新的陆地 ('1')，启动 BFS，表明发现了一个新的岛屿，计数器加 1。
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 四个方向
	count := 0

	// 广度优先搜索
	bfs := func(x, y int) {
		queue := [][2]int{{x, y}}
		for len(queue) > 0 {
			// 出队
			curr := queue[0]
			queue = queue[1:]
			for _, d := range directions {
				nx, ny := curr[0]+d[0], curr[1]+d[1]
				// 检查边界条件和是否为陆地
				if nx >= 0 && ny >= 0 && nx < m && ny < n && grid[nx][ny] == '1' {
					grid[nx][ny] = '0'                    // 标记为已访问
					queue = append(queue, [2]int{nx, ny}) // 加入队列
				}
			}
		}
	}

	// 遍历网格
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				grid[i][j] = '0' // 标记为已访问
				bfs(i, j)        // 启动 BFS
			}
		}
	}

	return count
}

// 上面的dfs实现方式太抽象,可读性太差,下面是简单易懂的
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
				grid[i][j] = '0' // 标记为已访问
				bfs(i, j)        // 启动 BFS
			}
		}
	}

	return count
}
