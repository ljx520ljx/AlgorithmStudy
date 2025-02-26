package Greedy

//135. 分发糖果
/*n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目。*/
/*
算法思路:贪心算法
遍历两次,分别只管右孩子评分高和左孩子评分高的情况
最后真正分的是这两次遍历所需要的糖果数取最大值
空间优化:
	rightCandy数组可以用一个变量代替
	只需要一个变量保存真实分的糖果总数
*/
func candy(ratings []int) int {
	n := len(ratings)
	leftCandy := make([]int, n) //从左往右,当右孩子比左孩子评分高,加糖果
	leftCandy[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			leftCandy[i] = leftCandy[i-1] + 1
		} else {
			leftCandy[i] = 1
		}
	}
	rightCandy := make([]int, n) //从右往左,当左孩子比右孩子评分高,加糖果
	rightCandy[n-1] = 1
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			rightCandy[i-1] = rightCandy[i] + 1
		} else {
			rightCandy[i-1] = 1
		}
	}
	realCandy := make([]int, n)
	res := 0
	for i := range realCandy {
		realCandy[i] = max(leftCandy[i], rightCandy[i])
		res += realCandy[i]
	}
	return res
}

func candy1(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}
