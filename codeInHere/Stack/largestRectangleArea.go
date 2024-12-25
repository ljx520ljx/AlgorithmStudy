package Stack

//84. 柱状图中最大的矩形
/*
算法思路:单调栈+常数优化
初始化：获取 heights 长度 n，创建 left、right 切片并初始化 right 元素为 n，创建空单调栈。
单调栈处理边界：遍历 heights，若栈非空且栈顶元素对应高度大于等于当前元素高度，更新栈顶元素右边界为 i 并出栈。
栈空时设 left[i] 为 -1，否则设为栈顶元素索引，再将 i 入栈。
计算最大面积：遍历heights，计算 (right[i] - left[i] - 1) * heights[i] 并更新最大面积。
*/

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	monoStack := []int{}
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}
