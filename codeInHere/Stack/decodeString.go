package Stack

// 394. 字符串解码
/*
算法思路:栈+迭代遍历
栈和迭代遍历：创建一个存储中间结果和数字的二维切片 stack 作为栈，初始化一个空字符串 curStr 用于存储当前正在构建的字符串，以及一个整数 num 用于存储当前数字。
数字处理：当遇到数字时，将其转换为整数 num。使用 strconv.Atoi 函数将数字字符转换为整数并累加到 num 上。
字母处理:当遇到字母时，将其添加到 curStr 中。
左括号处理:当遇到左括号时，将当前的 curStr 和 num 入栈，并将 curStr 重置为空字符串，将 num 重置为 0。
右括号处理:当遇到右括号时，从栈中弹出前一个状态，即 prevStr 和 prevNum，将当前的 curStr 重复 prevNum 次，并将结果与 prevStr 拼接，得到新的 curStr。
*/
import "strconv"

func decodeString(s string) string {
	var stack [][]interface{}
	curStr := ""
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit, _ := strconv.Atoi(string(s[i]))
			num = num*10 + digit
		} else if s[i] >= 'a' && s[i] <= 'z' {
			curStr += string(s[i])
		} else if s[i] == '[' {
			stack = append(stack, []interface{}{curStr, num})
			curStr = ""
			num = 0
		} else if s[i] == ']' {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevStr := prev[0].(string)
			prevNum := prev[1].(int)
			tmp := ""
			for j := 0; j < prevNum; j++ {
				tmp += curStr
			}
			curStr = prevStr + tmp
		}
	}
	return curStr
}
