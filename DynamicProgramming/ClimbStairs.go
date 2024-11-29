package DynamicProgramming

//爬楼梯
/*
算法思路:动态规划
每次可以爬 1 或 2 个台阶，假设爬到第n 阶的方法数是f(n)。
从第n−1阶爬1阶到n阶,或者从第n−2阶爬2阶到第n阶,因此状态转移方程为：
f(n)=f(n−1)+f(n−2)
*/

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// 使用动态规划法，初始化前两步
	prev1, prev2 := 1, 2

	// 从第三步开始计算到第 n 步的方法数
	for i := 3; i <= n; i++ {
		current := prev1 + prev2 // 当前方法数是前两步之和
		prev1 = prev2            // 更新 prev1 为 prev2
		prev2 = current          // 更新 prev2 为当前方法数
	}
	return prev2
}

// 自己写的递归实现,时间复杂度高,有大量重复子问题计算过程
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	var result int
	result = climbStairs(n-2) + climbStairs(n-1)
	return result
}

// 改进:使用记忆化递归（自顶向下）用一个数组存储已经计算过的结果，避免重复计算：
func climbStairs2(n int) int {
	memo := make([]int, n+1)
	return helper(n, memo)
}

func helper(n int, memo []int) int {
	if n <= 2 {
		return n
	}
	if memo[n] != 0 {
		return memo[n] // 如果已经计算过，直接返回
	}
	memo[n] = helper(n-1, memo) + helper(n-2, memo) // 计算并存储结果
	return memo[n]
}
