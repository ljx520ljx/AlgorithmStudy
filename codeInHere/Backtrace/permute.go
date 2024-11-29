package Backtrace

//全排列给定一个不含重复数字的数组nums,返回其所有可能的全排列.
/*
算法思路:回溯算法
利用递归和状态标记生成全排列。
在每一步中，通过选择一个尚未被使用的数字加入当前排列，然后递归生成剩下的排列。
递归完成后回溯，撤销选择以尝试新的排列路径。
*/

func permute(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	path := make([]int, n)
	onPath := make([]bool, n)
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			res = append(res, append([]int(nil), path...))
			return
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				backtrack(i + 1)
				onPath[j] = false
			}
		}
	}
	backtrack(0)
	return res
}
