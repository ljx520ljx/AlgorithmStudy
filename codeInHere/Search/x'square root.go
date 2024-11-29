package Search

//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//由于返回类型是整数，结果只保留整数部分 ，小数部分将被舍去 。
/*
算法思路:
初始化边界：
如果 x 小于 2，那么它的平方根就是自身，因此直接返回 x。
设置 left 为 1，right 为 x / 2。由于平方根不会超过 x / 2（对于 x >= 4），所以将 right 初始化为 x / 2。
二分查找：
使用 for left <= right 进行循环，计算中点 mid。
判断 mid * mid 是否等于 x，如果是，则直接返回 mid。
如果 mid * mid < x，说明平方根在右边，因此更新 left = mid + 1，并记录 result = mid（用于保留最近的小于 x 的整数）。
如果 mid * mid > x，说明平方根在左边，因此更新 right = mid - 1。
返回结果：
最终返回 result，即 x 的整数平方根。
*/
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	var result int

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
