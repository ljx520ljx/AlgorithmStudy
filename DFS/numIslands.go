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
