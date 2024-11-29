package Stack

//有效的括号给定一个只包括'(',')','{','}','[',']'的字符串s,判断字符串是否有效。
/*
算法思路:栈
利用栈来存储开放括号：
当遇到开放括号 (, [, { 时，将其压入栈中。
栈的特点是后进先出（LIFO），因此栈顶始终存储的是最新的未匹配的括号。
遇到闭合括号时进行匹配：
当遇到闭合括号 ), ], } 时，从栈中弹出栈顶元素，并检查是否与当前闭合括号匹配。
如果栈为空或者匹配失败，直接返回 false，因为此时括号的匹配顺序或数量不正确。
最终检查栈是否为空：
如果遍历结束时，栈仍然不为空，则说明存在未匹配的开放括号，应返回 false。
如果栈为空，则所有括号都正确匹配，返回 true。
*/

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := []rune{}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch) // 入栈
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' { // 检查栈顶是否匹配
				return false
			}
			stack = stack[:len(stack)-1] // 出栈
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0 // 如果栈为空，则括号完全匹配
}
