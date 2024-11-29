package Stack

// 最长有效括号
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
/*
算法思路:栈
左括号入栈
右括号时pop，栈非空时，当前下标与栈顶下标差即为有效括号长度;
栈空时，则push进当前下标充当左侧的哨兵节点
初始化栈时，push进−1当哨兵节点
*/

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return 0
	}
	var res int
	st := []int{-1}
	for i, j := range s {
		if j == '(' {
			st = append(st, i)
			continue
		}
		st = st[:len(st)-1]
		if len(st) == 0 {
			st = append(st, i)
		} else {
			res = max(res, i-st[len(st)-1])
		}
	}
	return res
}

// GPT写法
func longestValidParentheses1(s string) int {
	stack := []int{-1} // 初始化栈，-1 用于处理有效括号子串从索引 0 开始的情况
	maxLength := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i) // 遇到 '('，将其索引压入栈
		} else {
			stack = stack[:len(stack)-1] // 遇到 ')'，弹出栈顶元素，表示匹配的 '('
			if len(stack) > 0 {
				maxLength = max(maxLength, i-stack[len(stack)-1]) // 计算当前有效子串的长度
			} else {
				stack = append(stack, i) // 如果栈为空，说明没有匹配的 '('，将当前索引压入栈作为新的基准
			}
		}
	}

	return maxLength
}
