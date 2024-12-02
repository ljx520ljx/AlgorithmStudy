package Backtrace

//目标和:给你一个非负整数数组 nums 和一个整数 target 。
//向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
/*
算法思路:回溯
每个数字有两种选择：要么加上（+），要么减去（-）。
递归探索：对于数组中的每个数字，可以选择加上它或者减去它，并继续递归下去，直到遍历完所有数字。
基准条件：递归的终止条件是当数组中的所有数字都被遍历完后，检查当前的和是否等于目标值 target。如果相等，则说明找到了一种有效的组合。
回溯：每次选择加或减一个数字后，递归处理下一个数字，直到回到上一层进行不同的选择。
*/

func findTargetSumWays(nums []int, target int) (count int) {
	var backtrack func(int, int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	backtrack(0, 0)
	return
}
