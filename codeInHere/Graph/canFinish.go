package Graph

//课程表,即判断选修某些课程之前需要一些先修课程过程中是否存在环
/*
算法思路:拓扑排序
图的建模：可以把每个课程看作图中的一个节点，课程的先修关系可以看作有向边.例如[1,0]表示从课程0到课程1存在一条有向边,表示你需要先学习课程0再学习课程1。
判断环的存在：我们需要判断图中是否存在环。如果图中存在环，那么意味着你无法按照依赖关系完成课程的学习。判断环的方式可以使用深度优先搜索（DFS）或者拓扑排序。
拓扑排序的实现：
拓扑排序要求每次选取没有先修课程的节点（即没有入度的节点）。
使用入度数组来记录每个节点（课程）的入度（即当前课程依赖于多少课程）。
利用队列来管理当前可以学习的课程。
每次取出一个课程，并将所有依赖它的课程的入度减 1，如果某个课程的入度变为 0，则可以加入队列中继续学习。
如果最后队列中的课程数量等于 numCourses，说明所有课程都能学习；否则，说明存在环，无法完成所有课程。
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 创建一个入度数组，记录每个课程的入度
	inDegree := make([]int, numCourses)
	// 创建一个图，记录每个课程的依赖课程
	graph := make([][]int, numCourses)

	// 构建图和入度数组
	for _, prereq := range prerequisites {
		course := prereq[0]
		preCourse := prereq[1]
		graph[preCourse] = append(graph[preCourse], course)
		inDegree[course]++
	}

	// 创建一个队列，存储所有入度为0的课程
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 计数，记录已经完成学习的课程数量
	completedCourses := 0

	// 拓扑排序
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		completedCourses++

		// 遍历当前课程的依赖课程，减少它们的入度
		for _, nextCourse := range graph[course] {
			inDegree[nextCourse]--
			// 如果入度为0，加入队列
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// 如果完成的课程数量等于总课程数，返回true；否则返回false
	return completedCourses == numCourses
}

// 三色标记法
func canFinish1(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}

	colors := make([]int, numCourses)
	var dfs func(int) bool
	dfs = func(x int) bool {
		colors[x] = 1 // x 正在访问中
		for _, y := range g[x] {
			if colors[y] == 1 || colors[y] == 0 && dfs(y) {
				return true // 找到了环
			}
		}
		colors[x] = 2 // x 完全访问完毕
		return false  // 没有找到环
	}

	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false // 有环
		}
	}
	return true // 没有环
}
