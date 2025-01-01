package DFS

// 301. 删除无效的括号
/*
算法思路:dfs
1.计算多的需要删除的左括号或右括号数
2.dfs遍历字符进行剪枝、避免重复及根据情况删除括号并递归调用自身。
3.有效性检查函数 isValid,当计数器小于0或者不等于0,则判断为无效字符串。
*/

func removeInvalidParentheses(s string) (ans []string) {
	leftNum, rightNum := 0, 0
	// 计算leftNum, rightNum
	for _, j := range s {
		if j == '(' {
			leftNum++
		} else if j == ')' {
			if leftNum > 0 {
				leftNum--
			} else {
				rightNum++
			}
		}
	}
	dfs(&ans, s, 0, leftNum, rightNum)
	return
}

func isValid(str string) bool {
	cnt := 0
	for _, j := range str {
		if j == '(' {
			cnt++
		} else if j == ')' {
			cnt--
			// 说明当前出现了')'比'('多的情况
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func dfs(ans *[]string, str string, start, leftNum, rightNum int) {
	if leftNum == 0 && rightNum == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		// 去重,在当前循环中多的括号不会被删除,如果出现重复需要删除的括号不进行跳过,会导致答案添加了重复结果
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 剪枝,当需要删除的括号数大于剩余的字符数,则不需要进行dfs
		if leftNum+rightNum > len(str)-i {
			return
		}
		if leftNum > 0 && str[i] == '(' {
			dfs(ans, str[:i]+str[i+1:], i, leftNum-1, rightNum)
		}
		if rightNum > 0 && str[i] == ')' {
			dfs(ans, str[:i]+str[i+1:], i, leftNum, rightNum-1)
		}
	}
}
