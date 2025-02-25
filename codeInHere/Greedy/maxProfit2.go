package Greedy

//买卖股票的最佳时机,可以买卖多次
/*
算法思路:贪心算法
如果当前价格比前一天的价格高,就可以在前一天买入,在当前天卖出,从而获得利润
然后遍历整个价格数组,累加所有的利润,就是最大利润
*/
func maxProfit2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
