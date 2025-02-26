package Greedy

//45. 跳跃游戏 II
//从数组的第 0 个位置开始跳，跳的距离小于等于数组上对应的数。求出跳到最后一个位置需要的最短步数。
/*
算法思路:
	贪心算法
	每次在可跳范围内选择可以使得跳的更远的位置。
	每当跳到边界时，更新边界，并且步数加一,因为到边界了,说明一定选择了边界内的某一点作为一跳.
	加上题目保证了一定会到达最后一个位置,所以不需要考虑到底能不能到达最后一个位置.
*/

func jump(nums []int) int {
	currRight := 0
	maxRight := 0
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		maxRight = max(maxRight, i+nums[i])
		if i == currRight {
			currRight = maxRight
			res++
		}
	}
	return res
}
