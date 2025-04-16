package String

// 227. 基本计算器 II
// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。整数除法仅保留整数部分。
/*
算法思路: 栈 + 运算符优先级
1. 初始化一个栈 `stack` 用于存储待加减的数字（乘除法会被立即计算）。
2. 初始化 `preSign` 为 '+'，表示数字前面的符号。
3. 初始化 `num` 为 0，用于累积当前解析的数字。
4. 遍历字符串 `s`：
   - 如果当前字符 `ch` 是数字，则更新 `num = num*10 + int(ch-'0')`。
   - 如果当前字符 `ch` 不是数字且不是空格，或者到达字符串末尾 (i == len(s)-1)，说明一个数字解析完成。
   - 此时根据 `preSign` 的值执行相应的操作：
     - `+`: 将 `num` 压入栈。
     - `-`: 将 `-num` 压入栈。
     - `*`: 从栈顶取出元素，与 `num` 相乘，结果再压回栈顶。
     - `/`: 从栈顶取出元素，与 `num` 相除（整数除法），结果再压回栈顶。
   - 处理完数字后，更新 `preSign` 为当前字符 `ch`，并将 `num` 重置为 0。
5. 遍历结束后，栈中剩下的所有元素都是需要相加的（减法已通过负数处理），将它们累加得到最终结果 `ans`。
*/
func calculate(s string) (ans int) {
	// 使用切片模拟栈，存储待加减的运算数
	stack := []int{}
	// preSign 记录当前数字 num 前面的那个运算符
	preSign := '+'
	// num 记录当前累积的数字
	num := 0
	// 遍历字符串
	for i, ch := range s {
		// 判断当前字符是否是数字
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			// 如果是数字，累加到 num
			num = num*10 + int(ch-'0')
		}
		// 判断是否需要处理 num：
		// 1. 当前字符不是数字且不是空格
		// 2. 或者已经到达字符串的最后一个字符 (必须处理最后一个数字)
		if !isDigit && ch != ' ' || i == len(s)-1 {
			// 根据 num 前面的运算符 preSign 来决定如何处理 num
			switch preSign {
			case '+':
				// 加法：直接入栈
				stack = append(stack, num)
			case '-':
				// 减法：将相反数入栈
				stack = append(stack, -num)
			case '*':
				// 乘法：取出栈顶元素，与 num 相乘后，结果再入栈
				stack[len(stack)-1] *= num
			default: // 除法
				// 除法：取出栈顶元素，与 num 相除后，结果再入栈
				stack[len(stack)-1] /= num
			}
			// 更新 preSign 为当前遇到的运算符
			preSign = ch
			// 重置 num，准备累积下一个数字
			num = 0
		}
	}
	// 遍历栈，将所有元素累加，得到最终结果
	for _, v := range stack {
		ans += v
	}
	// 返回结果
	return
}
