package SpecialAlgorithm

//螺旋矩阵:给你一个m行n列的矩阵 matrix
//请按照顺时针螺旋顺序,即螺旋内缩,返回矩阵中的所有元素。
/*
算法思路:
边界定义：用 top、bottom、left 和 right 定义当前需要遍历的矩阵边界。
四个方向的遍历：
从左到右：固定上边界 top，从左到右遍历当前行。
从上到下：固定右边界 right，从上到下遍历当前列。
从右到左：固定下边界 bottom，从右到左遍历当前行。
从下到上：固定左边界 left，从下到上遍历当前列。
边界收缩：每次完成一轮遍历后，收缩对应的边界（top++、bottom--、left++、right--）。
终止条件：当上下边界或左右边界交错时，遍历结束。
*/

func spiralOrder(matrix [][]int) []int {
	// 初始化结果数组
	var result []int

	if len(matrix) == 0 {
		return result
	}

	// 定义四个边界
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// 从左到右
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++ // 移动上边界

		// 从上到下
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right-- // 移动右边界

		// 判断是否仍在范围内
		if top <= bottom {
			// 从右到左
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom-- // 移动下边界
		}

		// 判断是否仍在范围内
		if left <= right {
			// 从下到上
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++ // 移动左边界
		}
	}

	return result
}
