package Backtrace

import "fmt"

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
		//这里index == len(nums)和sum == target要分开判断
		//不能写成index == len(nums)&&sum == target,因为一旦index == len(nums)就要return,和sum == target无关
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

func findTargetSumWays1(nums []int, target int) int {
	n := len(nums)
	var res int
	var backtrack func(int, int)
	backtrack = func(i, target int) {
		if i == n {
			if target == 0 {
				res++
			}
			return
		}
		backtrack(i+1, target-nums[i])
		backtrack(i+1, target+nums[i])
	}
	backtrack(0, target)
	return res
}

// 上面的回溯没有记忆优化,存在大量子问题的重复计算,下面的回溯采用记忆优化
// 再加上剪枝优化
func findTargetSumWays2(nums []int, target int) int {
	n := len(nums)
	// 计算 nums 所有元素的总和
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// 如果 target 超出了可能的最大范围，直接返回 0
	if target > sum || target < -sum {
		return 0
	}

	// 使用一个map来存储已计算的结果
	memo := make(map[string]int)

	var backtrack func(int, int) int
	backtrack = func(i, target int) int {
		// 如果已经到达数组末尾
		if i == n {
			if target == 0 {
				return 1 // 找到了一个符合条件的解
			}
			return 0 // 没有符合条件的解
		}

		// 使用字符串拼接索引和目标值作为哈希表的key
		key := fmt.Sprintf("%d,%d", i, target)
		// 如果当前状态已经计算过，直接返回记忆化的结果
		if val, ok := memo[key]; ok {
			return val
		}

		// 剪枝：如果当前目标值已经超出剩余元素和的范围，则返回 0
		if target > sum || target < -sum {
			return 0
		}

		// 递归计算两种选择：+nums[i] 和 -nums[i]
		res := backtrack(i+1, target-nums[i]) + backtrack(i+1, target+nums[i])

		// 将计算结果存储到memo中
		memo[key] = res

		return res
	}

	return backtrack(0, target)
}
