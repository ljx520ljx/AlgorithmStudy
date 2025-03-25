package Backtrace

//复原IP地址
//给定一个只包含数字的字符串s，用以表示一个IP地址，返回所有可能的有效IP地址，这些地址可以通过在s中插入'.'来形成。
//有效IP地址正好由四个整数（每个整数位于0到255之间组成，且不能含有前导0），整数之间用'.'分隔。
//例如："0.1.2.201" 和 "192.168.1.1" 是有效IP地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是无效IP地址。
/*
算法思路:回溯算法
1. 边界条件处理：
   - 如果字符串长度小于4或大于12，直接返回空结果（因为IP地址必须是4段，每段1-3个数字）
2. 回溯算法实现：
   - 使用回溯函数递归处理每一段IP地址
   - 每一段可以取1-3个字符
   - 对于每一段需要检查：
     * 数字范围是否在0-255之间
     * 是否有前导0
3. 终止条件：
   - 当已经形成4段且使用完所有字符时，将当前组合加入结果
   - 当已经形成4段但还有剩余字符，或字符用完但还不足4段时，直接返回
4. 特殊情况处理：
   - 对于数字0，只能单独成为一段
   - 其他数字不能有前导0
*/

import "strconv"

func restoreIpAddresses(s string) []string {
	var result []string
	if len(s) < 4 || len(s) > 12 {
		return result
	}

	var backtrack func(start int, segments []string)
	backtrack = func(start int, segments []string) {
		// 如果已经有4段且用完所有字符，则找到一个有效解
		if len(segments) == 4 && start == len(s) {
			result = append(result, segments[0]+"."+segments[1]+"."+segments[2]+"."+segments[3])
			return
		}

		// 如果已经有4段但还有剩余字符，或者字符用完但还不足4段，则返回
		if len(segments) == 4 || start == len(s) {
			return
		}

		// 处理前导0的特殊情况
		if s[start] == '0' {
			segments = append(segments, "0")
			backtrack(start+1, segments)
			return
		}

		// 尝试当前位置开始的1到3个字符
		curr := ""
		for i := start; i < start+3 && i < len(s); i++ {
			curr += string(s[i])
			if num, _ := strconv.Atoi(curr); num <= 255 {
				segments = append(segments, curr)
				backtrack(i+1, segments)
				segments = segments[:len(segments)-1]
			}
		}
	}

	backtrack(0, []string{})
	return result
}
