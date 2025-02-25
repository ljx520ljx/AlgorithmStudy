package DynamicProgramming

import "math"

//买卖股票的最佳时机,只能买卖一次
/*
算法思路:动态规划
遍历价格数组时，minPrice初始化为一个很大的值.
记录当前的最低价格minPrice,若比minPrice价格更低,更新minPrice.
对于每个价格，计算当前价格卖出的潜在利润，即 当前价格 - minPrice。
用一个变量 maxProfit 存储利润中的最大值,随时更新maxProfit
最终的 maxProfit 就是答案。
*/

func maxProfit(prices []int) int {
	maxProfit := 0
	if len(prices) == 1 {
		return maxProfit
	}
	minPrice := math.MaxInt32

	for i := 0; i < len(prices); i++ {
		if prices[i] <= minPrice {
			minPrice = prices[i]
		} else if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
