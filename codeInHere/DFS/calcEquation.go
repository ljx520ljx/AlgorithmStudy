package DFS

//399. 除法求值
/*
算法思路:dfs深度优先遍历
该算法使用图的思想，将变量视为图中的节点，方程视为边。
对于每个方程 a/b = value，在图中添加两条边：a 到 b 的边权重为 value，b 到 a 的边权重为 1/value。
对于每个查询 [start, end]，使用深度优先搜索 dfs 从 start 开始搜索到 end 的路径，将路径上的边权重相乘得到结果。
如果在 dfs 过程中无法到达 end 或 start 不在图中，返回 -1.0。
*/

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 构建图的邻接表存储变量之间的关系
	graph := make(map[string]map[string]float64)
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		if _, exists := graph[a]; !exists {
			graph[a] = make(map[string]float64)
		}
		if _, exists := graph[b]; !exists {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1 / values[i]
	}
	// 深度优先搜索
	var dfs func(start, end string, visited map[string]bool) float64
	dfs = func(start, end string, visited map[string]bool) float64 {
		// 检查该节点是否存在
		if _, exists := graph[start]; !exists {
			return -1.0
		}
		// 检查是不是自己,是自己就返回1
		if start == end {
			return 1.0
		}
		// 标记已经访问过
		visited[start] = true
		// 遍历邻居节点
		for neighbor, value := range graph[start] {
			if !visited[neighbor] {
				// 递归搜索邻居节点
				res := dfs(neighbor, end, visited)
				// 如果找到路径，返回结果
				if res != -1.0 {
					return value * res
				}
			}
		}
		// 未找到路径
		return -1.0
	}

	result := make([]float64, len(queries))
	// 遍历查询储存答案
	for i, query := range queries {
		// 初始化visited
		visited := make(map[string]bool)
		result[i] = dfs(query[0], query[1], visited)
	}
	return result
}
