package DynamicProgramming

//85. 最大矩形
/*
算法思路:
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var res int
	height := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := range height {
			height[j] = (height[j] + 1) * int(matrix[i][j]-'0')
			minheight := height[j]
			for k := j; k >= 0 && height[k] > 0; k-- {
				minheight = min(minheight, height[k])
				res = max(res, minheight*(j-k+1))
			}
		}
	}

	return res
}
