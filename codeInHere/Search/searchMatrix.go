package Search

// 240. 搜索二维矩阵 II
/*
算法思路:暴力直接查找+剪枝
两层循环遍历暴力查找
剪枝:因为二维数组的每行每列都是升序排列的,所以当target小于当前元素时,就可以跳出当前循环进入下一层循环
*/
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if target < matrix[i][j] {
				break
			} else if target == matrix[i][j] {
				return true
			}
		}
	}
	return false
}

/*
算法思路: 从右上角开始的"缩减搜索空间"
1. 观察矩阵特性：
  - 每行元素从左到右升序排列
  - 每列元素从上到下升序排列
  - 这形成了一个特殊的排序结构，可以选择一个特殊的起点进行搜索

2. 搜索策略：从矩阵的右上角(0, n-1)开始搜索
  - 这个位置的特殊性在于：向左元素变小，向下元素变大
  - 这使我们每次比较后可以排除一行或一列

3. 搜索过程：
  - 比较当前元素与目标值:
  - 如果相等，找到目标，返回true
  - 如果当前元素 > 目标值，说明当前列所有元素都大于目标值（因为下方元素更大），因此向左移动(y--)
  - 如果当前元素 < 目标值，说明当前行所有元素都小于目标值（因为左侧元素更小），因此向下移动(x++)
  - 重复以上过程，直到找到目标或搜索区域为空（x越界或y越界）

4. 终止条件：
  - 找到目标值，返回true
  - 指针超出矩阵边界(x >= m 或 y < 0)，返回false
*/
func searchMatrix1(matrix [][]int, target int) bool {
	// 获取矩阵尺寸
	m, n := len(matrix), len(matrix[0])
	// 初始化指针位置为右上角
	x, y := 0, n-1

	// 搜索过程：当指针在矩阵范围内时继续
	for x < m && y >= 0 {
		// 找到目标
		if matrix[x][y] == target {
			return true
		}
		// 当前值大于目标值，向左移动（减小当前值）
		if matrix[x][y] > target {
			y--
		} else { // 当前值小于目标值，向下移动（增大当前值）
			x++
		}
	}
	// 搜索结束未找到目标
	return false
}
