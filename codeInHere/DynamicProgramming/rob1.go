package DynamicProgramming

// 213. 打家劫舍 II
//打家劫舍1的基础上，房屋是环形的，即第一个和最后一个房屋相邻。
/*
算法思路:动态规划
在打家劫舍的基础上,分两种情况:
1.偷第一家,不偷最后一家rob(nums[0:n-1])
2.不偷第一家,偷最后一家rob(nums[1:n])
取两种情况的最大值
*/
func rob1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(rob(nums[1:n]), rob(nums[0:n-1]))
	//return max(rob1(nums[2:n-1])+nums[0],rob1(nums[1:n]))
}
