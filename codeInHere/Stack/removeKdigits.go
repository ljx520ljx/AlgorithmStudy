package Stack

import "strings"

// 402. 移掉 K 位数字
// 给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。
/*
算法思路: 贪心 + 单调栈
1. 核心思想：为了使结果数字最小，我们希望高位（左边）的数字尽可能小。因此，我们从左到右遍历数字，如果发现一个数字比它前面的数字小，那么就应该移除它前面的那个较大的数字（因为它在高位，移除它带来的收益最大）。
2. 使用栈（Go中用切片模拟）：维护一个单调递增（或非递减）的栈。这个栈存储的是我们当前构建的最小数字的候选字符。
3. 遍历输入字符串 `num`：
   - 对于当前数字 `digit`：
     a. **检查栈顶**: 当栈不为空、`k > 0` 并且栈顶元素 `stack[top]` 大于当前数字 `digit` 时：
        - 说明栈顶元素是一个“峰值”，移除它可以让更高位的数字变小。
        - 将栈顶元素弹出 `stack = stack[:len(stack)-1]`。
        - 消耗一次移除机会 `k--`。
        - **重复**这个检查，直到栈顶元素不大于当前数字，或者栈为空，或者 `k` 耗尽。
     b. **当前数字入栈**: 将当前数字 `digit` 压入栈中 `stack = append(stack, digit)`。
4. 处理剩余的 `k`：
   - 遍历完所有数字后，如果 `k` 仍然大于 0，说明原始数字本身就是单调递增或非递减的（或者移除次数不足）。
   - 在这种情况下，为了使数字最小，我们应该从**末尾**（即栈顶）移除剩余的 `k` 个数字（因为末尾的数字权重最低）。
   - `stack = stack[:len(stack)-k]`。
5. 构建结果字符串：
   - 栈中现在存储的就是移除 k 位后剩下的数字字符。
   - 将栈中的字符连接成字符串。
6. 处理前导零：
   - 使用 `strings.TrimLeft(result, "0")` 移除结果字符串的前导零。
7. 处理特殊情况：
   - 如果移除前导零后字符串为空（说明原数字移除k位后只剩下0，或者原数字就是"0"且k>0），则返回 "0"。
   - 否则返回处理后的字符串。
*/
func removeKdigits(num string, k int) string {
	// 确保不会移除超过栈长度的元素
	if k >= len(num) {
		return "0" // 如果移除次数大于等于栈长度，结果为 "0"
	}
	// 使用 byte 切片模拟单调栈
	var stack []byte
	// 遍历数字字符串
	for i := 0; i < len(num); i++ {
		digit := num[i] // 当前数字字符 (byte)

		// 当栈不为空，k > 0，且栈顶元素 > 当前数字时，移除栈顶（贪心）
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > digit {
			stack = stack[:len(stack)-1] // 弹出栈顶
			k--                          // 消耗一次移除机会
		}
		// 将当前数字压入栈
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	//也可以用这个方式处理前导零
	//index:=0
	//for index<len(stack){
	//	if stack[index]=='0'{
	//		index++
	//	}else{
	//		break
	//	}
	//}
	//res:=String(stack[index:])
	// 将栈转换成字符串
	result := string(stack)
	// 处理前导零
	result = strings.TrimLeft(result, "0")
	// 如果移除前导零后字符串为空，说明结果是 0
	if result == "" {
		return "0"
	}
	return result
}
