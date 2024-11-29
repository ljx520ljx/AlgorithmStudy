package Backtrace

//括号生成
//数字n代表生成括号的对数,请你设计一个函数,用于能够生成所有可能的并且有效的括号组合.
/*
算法思路:回溯
递归和回溯：使用回溯法生成所有可能的括号组合，同时保证生成的括号是有效的。
一个有效的括号组合需要满足：右括号数量不超过左括号数量，并且左括号数量不超过 n。
递归状态：
current: 当前生成的括号字符串。
open: 当前字符串中左括号的数量。
close: 当前字符串中右括号的数量。
递归结束条件：
当 current 的长度等于 2*n 时，表示已经生成了一个完整的括号组合，添加到结果列表中。
递归过程：
如果左括号数量小于 n，可以继续添加左括号。
如果右括号数量小于左括号数量，可以继续添加右括号。
效率：
每次递归都根据规则生成有效的部分括号组合，避免生成无效的组合，提升效率。
*/

func generateParenthesis(n int) []string {
	var result []string
	// 定义递归函数
	var backtrack func(current string, open, close int)
	backtrack = func(current string, open, close int) {
		// 如果当前字符串的长度达到了 2*n，加入结果
		if len(current) == 2*n {
			result = append(result, current)
			return
		}
		// 如果可以添加左括号
		if open < n {
			backtrack(current+"(", open+1, close)
		}
		// 如果可以添加右括号
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}
	// 从空字符串开始
	backtrack("", 0, 0)
	return result
}
