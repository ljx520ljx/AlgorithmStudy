package string

import (
	"strconv"
	"strings"
)

//43. 字符串相乘
//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
/*
算法思路: 模拟竖式乘法
1. 特殊处理：如果num1或num2是"0"，直接返回"0"。
2. 创建结果数组：结果字符串的最大长度不会超过 num1.len + num2.len。创建一个长度为 m+n 的整数数组 `res` 并初始化为0，用于存储竖式乘法每位的计算结果。
3. 逐位相乘：
   - 从右到左遍历 num1 (索引 i)。
   - 从右到左遍历 num2 (索引 j)。
   - 取出当前位的数字 `digit1 = num1[i] - '0'` 和 `digit2 = num2[j] - '0'`。
   - 计算它们的乘积 `product = digit1 * digit2`。
   - 确定乘积在 `res` 数组中的位置：
     - 低位 (个位) 在 `res[i + j + 1]`。
     - 高位 (十位，即进位) 在 `res[i + j]`。
   - 更新 `res` 数组：
     - `currentSum = res[i + j + 1] + product` (当前位置的原有值加上新乘积)。
     - `res[i + j + 1] = currentSum % 10` (更新低位的值)。
     - `res[i + j] += currentSum / 10` (将进位加到高位)。
4. 处理结果数组：
   - `res` 数组现在存储了乘积的每一位数字，但可能有前导零。
   - 找到第一个非零数字的索引 `start`。
   - 如果整个数组都是零（说明结果是0），返回 "0"。
5. 构建结果字符串：
   - 从 `start` 索引开始遍历 `res` 数组。
   - 将每个数字转换为字符串并拼接到结果字符串中。
6. 返回结果字符串。
*/
func multiply(num1 string, num2 string) string {
	// 1. 特殊处理：如果任一数字是 "0"，结果就是 "0"
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	// 2. 创建结果数组，长度为 m+n，存储每一位的计算结果
	res := make([]int, m+n)

	// 3. 逐位相乘，模拟竖式乘法
	// 从 num1 的末尾开始遍历
	for i := m - 1; i >= 0; i-- {
		// 获取 num1 当前位的数字
		digit1 := int(num1[i] - '0')
		// 从 num2 的末尾开始遍历
		for j := n - 1; j >= 0; j-- {
			// 获取 num2 当前位的数字
			digit2 := int(num2[j] - '0')
			// 计算当前两位数字的乘积
			product := digit1 * digit2

			// 确定乘积在结果数组中的位置
			pos1 := i + j + 1 // 低位位置
			pos2 := i + j     // 高位位置

			// 将乘积加到结果数组的对应位置上
			// currentSum 包含当前位置的原有值和新的乘积
			currentSum := res[pos1] + product
			// 更新低位的值（只保留个位）
			res[pos1] = currentSum % 10
			// 将进位加到高位
			res[pos2] += currentSum / 10
		}
	}

	// 4. 处理结果数组，去除前导零
	start := 0
	// 找到第一个非零数字的索引
	for start < len(res)-1 && res[start] == 0 {
		start++
	}
	// 5. 构建结果字符串
	var sb strings.Builder
	// 从第一个非零数字开始遍历
	for i := start; i < len(res); i++ {
		// 将数字转换为字符串并追加
		sb.WriteString(strconv.Itoa(res[i]))
	}

	// 6. 返回结果字符串
	return sb.String()
}
