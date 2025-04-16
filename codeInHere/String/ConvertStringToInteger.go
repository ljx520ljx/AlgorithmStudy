package String

//把字符串转换为整数
/*
算法思路:
忽略前导空格：我们需要跳过字符串开头的所有空格字符。
处理符号：检测第一个非空格字符是否是 '+' 或 '-'，来确定数值的正负。如果没有符号，默认为正数。
读取数字：逐个字符读取，直到遇到非数字字符或到达字符串末尾。忽略任何前导的零。
范围检查：将结果限制在 32 位有符号整数的范围内，即 [-2^31, 2^31 - 1]，如果超出范围，返回相应的边界值。
返回结果：返回转换后的整数。
*/
import "math"

func myAtoi(s string) int {
	// 初始化索引、符号和结果
	i, sign, result := 0, 1, 0
	n := len(s)

	// 1. 跳过前导空格
	for i < n && s[i] == ' ' {
		i++
	}

	// 2. 处理符号
	if i < n {
		if s[i] == '+' {
			sign = 1
			i++
		} else if s[i] == '-' {
			sign = -1
			i++
		}
	}

	// 3. 读取数字
	for i < n && s[i] >= '0' && s[i] <= '9' {
		// 将当前字符转换为数字
		digit := int(s[i] - '0')

		// 4. 检查是否超出 32 位整数范围
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		// 更新结果
		result = result*10 + digit
		i++
	}

	// 5. 返回结果
	return result * sign
}
