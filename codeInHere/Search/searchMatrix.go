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
