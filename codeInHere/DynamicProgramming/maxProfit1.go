package DynamicProgramming

// 122. 买卖股票的最佳时机含冷冻期
/*
算法思路:动态规划
创建两个切片buy,sell.buy[i]表示第i天持有股票的最大收益,sell[i]表示第i天不持有股票的最大收益
初始化buy[0]=-prices[0],sell[0]=0
第i天持有股票的最大收益=max(第i-1天持有股票的最大收益,第i-2天不持有股票的最大收益-当天股票的价格)因为有冷冻期.
第i天不持有股票的最大收益=max(第i-1天不持有股票的最大收益,第i-1天持有股票的最大收益+当天股票的价格)
*/

func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := make([]int, n)
	sell := make([]int, n)
	//初始化
	buy[0] = -prices[0]
	sell[0] = 0
	if n >= 2 {
		buy[1] = max(-prices[0], -prices[1])
		sell[1] = max(0, prices[1]-prices[0])
	}
	for i := 2; i < n; i++ {
		//第i天持有股票的最大收益=max(第i-1天持有股票的最大收益,第i-2天不持有股票的最大收益-当天股票的价格)因为有冷冻期.
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		//第i天不持有股票的最大收益=max(第i-1天不持有股票的最大收益,第i-1天持有股票的最大收益+当天股票的价格)
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[n-1]
}
